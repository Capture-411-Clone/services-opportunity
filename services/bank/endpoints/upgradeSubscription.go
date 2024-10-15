package endpoints

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/subscription"
)

type UpgradeSubscriptionRequestBody struct {
	RedirectPath string `json:"redirect_path"`
	NewPriceID   string `json:"new_price_id"`
	NewProductID string `json:"new_product_id"`
}

func (data *UpgradeSubscriptionRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"new_price_id":   gocriteria.New("new_price_id").Required(),
		"new_product_id": gocriteria.New("new_product_id").Required(),
		"redirect_path":  gocriteria.New("redirect_path").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"new_price_id":   "New Price ID",
			"new_product_id": "New Product ID",
			"redirect_path":  "Redirect Path",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) UpgradeSubscription(ctx context.Context, input UpgradeSubscriptionRequestBody) (UpgradeSubscriptionResponseData, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		return UpgradeSubscriptionResponseData{}, response.ErrorNotFound("User not found")
	}

	cust, err := customer.Get(user.StripeCustomerID, nil)
	if err != nil {
		// Handle error (e.g., customer not found)
		return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("Error retrieving customer")
	}

	if user.SubscriptionID == "" {
		return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("User does not have a subscription")
	}

	// Retrieve the current subscription to get its items, assuming it has only one item.
	currentSub, err := subscription.Get(user.SubscriptionID, nil)
	if err != nil {
		s.logger.Errorf("Error retrieving subscription: %v\n", err)
		return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("Error retrieving subscription")
	}

	currentPriceAmount := currentSub.Items.Data[0].Price.UnitAmount

	// get new price unit amount
	newPrice, err := price.Get(input.NewPriceID, nil)
	if err != nil {
		s.logger.With(ctx).Errorf("price.Get: %v", err)
		return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	newPriceAmount := newPrice.UnitAmount

	if newPriceAmount <= currentPriceAmount {
		return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("You can't upgrade to a price that is less than or equal to the current price.")
	}

	var ses *stripe.CheckoutSession
	var newSub *stripe.Subscription
	// 2. Create a new subscription with the new price ID
	if cust.InvoiceSettings.DefaultPaymentMethod == nil {
		domain := config.AppConfig.StripeRedirectDomain

		basePath, err := url.JoinPath(domain, input.RedirectPath)
		if err != nil {
			s.logger.With(ctx).Errorf("url.JoinPath: %v", err)
			return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
		}

		checkoutParams := &stripe.CheckoutSessionParams{
			Mode:     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
			Customer: stripe.String(cust.ID),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					Price:    stripe.String(input.NewPriceID),
					Quantity: stripe.Int64(1),
				},
			},
			SuccessURL: stripe.String(basePath + "?success=true&session_id={CHECKOUT_SESSION_ID}"),
			CancelURL:  stripe.String(basePath + "?canceled=true"),
			// NOTE: This is not needed cuz updating means these are already existed!
			// AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			// 	Enabled: stripe.Bool(true),
			// },
			// CustomerUpdate: &stripe.CheckoutSessionCustomerUpdateParams{
			// 	Address: stripe.String("auto"),
			// },
		}

		ses, err = session.New(checkoutParams)
		if err != nil {
			s.logger.With(ctx).Errorf("session.New: %v", err)
			return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("Could not proceed with payment proccess")
		}
	} else {
		params := &stripe.SubscriptionParams{
			Customer: stripe.String(user.StripeCustomerID),
			Items: []*stripe.SubscriptionItemsParams{
				{
					Price:    stripe.String(input.NewPriceID),
					Quantity: stripe.Int64(1),
				},
			},
			ProrationBehavior: stripe.String("none"), // This ensures no proration calculation
			AutomaticTax: &stripe.SubscriptionAutomaticTaxParams{
				Enabled: stripe.Bool(true),
			},
		}
		var createErr error
		newSub, createErr = subscription.New(params)
		if createErr != nil {
			s.logger.Errorf("Error creating new subscription: %v\n", createErr)
			return UpgradeSubscriptionResponseData{}, response.ErrorBadRequest("Error creating new subscription")
		}
	}

	return UpgradeSubscriptionResponseData{
		SessionID:    ses.ID,
		SessionURL:   ses.URL,
		Subscription: newSub,
	}, response.ErrorResponse{}
}

type UpgradeSubscriptionResponseData struct {
	SessionID    string               `json:"session_id" openapi:"example:cs_test_1234567890;"`
	SessionURL   string               `json:"session_url" openapi:"example:https://checkout.stripe.com/pay/cs_test_1234567890;"`
	Subscription *stripe.Subscription `json:"subscription"`
}
