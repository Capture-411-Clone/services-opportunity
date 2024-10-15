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

type CreateSubscriptionCheckoutSessionRequestBody struct {
	LookupKey    string `json:"lookup_key"`
	RedirectPath string `json:"redirect_path"`
}

func (data *CreateSubscriptionCheckoutSessionRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"lookup_key":    gocriteria.New("lookup_key").MinLength(3).Required(),
		"redirect_path": gocriteria.New("redirect_path").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"lookup_key":    "Lookup Key",
			"redirect_path": "Redirect Path",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) CreateSubscriptionCheckoutSession(ctx context.Context, input CreateSubscriptionCheckoutSessionRequestBody) (string, string, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateSubscriptionCheckout: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	var customerID string
	if user.StripeCustomerID == "" {
		// Create a new Stripe customer because one does not exist
		customerParams := &stripe.CustomerParams{
			Email: stripe.String(user.Email.String),
		}
		cust, err := customer.New(customerParams)
		if err != nil {
			s.logger.With(ctx).Errorf("customer.New: %v", err)
			return "", "", response.ErrorBadRequest(nil, "Could not create Stripe customer")
		}
		customerID = cust.ID

		// Store the new Stripe Customer ID in your database
		user.StripeCustomerID = customerID
		if err := s.db.Save(&user).Error; err != nil {
			s.logger.With(ctx).Errorf("Save StripeCustomerID: %v", err)
			return "", "", response.ErrorBadRequest(nil, "Could not update user with Stripe Customer ID")
		}
	} else {
		customerID = user.StripeCustomerID
	}

	subparams := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
		Status:   stripe.String(string(stripe.SubscriptionStatusActive)),
	}
	subparams.AddExpand("data.default_payment_method")
	subi := subscription.List(subparams)

	// If the customer has active subscriptions, return an error
	if subi.Next() {
		return "", "", response.ErrorBadRequest(nil, "You already have an active subscription. Please upgrade or downgrade your plan as needed.")
	}

	params := &stripe.PriceListParams{
		LookupKeys: stripe.StringSlice([]string{
			input.LookupKey,
		}),
	}

	i := price.List(params)

	var price *stripe.Price
	for i.Next() {
		p := i.Price()
		price = p
	}

	domain := config.AppConfig.StripeRedirectDomain

	basePath, err := url.JoinPath(domain, input.RedirectPath)
	if err != nil {
		s.logger.With(ctx).Errorf("url.JoinPath: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	checkoutParams := &stripe.CheckoutSessionParams{
		Mode:     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Customer: stripe.String(customerID),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(price.ID),
				Quantity: stripe.Int64(1),
			},
		},
		// TODO: add a path where user was and get it from front end initially
		SuccessURL: stripe.String(basePath + "?success=true&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(basePath + "?canceled=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
		CustomerUpdate: &stripe.CheckoutSessionCustomerUpdateParams{
			Address: stripe.String("auto"),
		},
		// Discounts: []*stripe.CheckoutSessionDiscountParams{
		// 	{Coupon: stripe.String("DUBCRAFT2021")},
		// },
		AllowPromotionCodes: stripe.Bool(true),
	}

	ses, err := session.New(checkoutParams)

	if err != nil {
		s.logger.With(ctx).Errorf("session.New: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	return ses.ID, ses.URL, response.ErrorResponse{}
}

type CreateSubscriptionCheckoutSessionResponseData struct {
	SessionID  string `json:"session_id" openapi:"example:cs_test_1234567890;"`
	SessionURL string `json:"session_url" openapi:"example:https://checkout.stripe.com/pay/cs_test_1234567890;"`
}
