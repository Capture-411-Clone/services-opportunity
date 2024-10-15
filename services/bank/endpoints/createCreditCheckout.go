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
	"github.com/stripe/stripe-go/v76/product"
	"github.com/stripe/stripe-go/v76/subscription"
)

type CreateCreditCheckoutSessionRequestBody struct {
	Amount       int    `json:"amount"`
	RedirectPath string `json:"redirect_path"`
}

func (data *CreateCreditCheckoutSessionRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"amount":        gocriteria.New("amount").Required(),
		"redirect_path": gocriteria.New("redirect_path").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"amount":        "Amount",
			"redirect_path": "Redirect Path",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) CreateCreditCheckoutSession(ctx context.Context, input CreateCreditCheckoutSessionRequestBody) (string, string, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateCreditCheckout: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
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

	domain := config.AppConfig.StripeRedirectDomain

	subscriptionID := user.SubscriptionID
	subscription, err := subscription.Get(subscriptionID, nil)
	if err != nil {
		s.logger.With(ctx).Errorf("subscription.Get: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	product, err := product.Get(subscription.Items.Data[0].Price.Product.ID, nil)
	if err != nil {
		s.logger.With(ctx).Errorf("product.Get: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	creditProductID := product.Metadata["credit_product_id"]
	if creditProductID == "" {
		s.logger.With(ctx).Error("credit_product_id not found in product metadata")
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	priceParams := &stripe.PriceListParams{
		Product: stripe.String(creditProductID),
	}
	priceList := price.List(priceParams)
	var creditPrice *stripe.Price
	for priceList.Next() {
		p := priceList.Price()
		// Assuming you want to select the first price or a specific one based on your logic
		creditPrice = p
		break // or specific logic to select the correct price
	}
	if creditPrice == nil {
		return "", "", response.ErrorBadRequest(nil, "Could not find price for credit product")
	}

	basePath, err := url.JoinPath(domain, input.RedirectPath)
	if err != nil {
		s.logger.With(ctx).Errorf("url.JoinPath: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	// Adjust checkoutParams for one-time purchase
	checkoutParams := &stripe.CheckoutSessionParams{
		Mode:     stripe.String(string(stripe.CheckoutSessionModePayment)),
		Customer: stripe.String(customerID),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(creditPrice.ID),
				Quantity: stripe.Int64(int64(input.Amount)),
			},
		},
		SuccessURL: stripe.String(basePath + "?success=true&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(basePath + "?canceled=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
		CustomerUpdate: &stripe.CheckoutSessionCustomerUpdateParams{
			Address: stripe.String("auto"),
		},
		// Adding invoice creation parameters
		InvoiceCreation: &stripe.CheckoutSessionInvoiceCreationParams{
			Enabled: stripe.Bool(true),
		},
	}

	ses, err := session.New(checkoutParams)

	if err != nil {
		s.logger.With(ctx).Errorf("session.New: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	return ses.ID, ses.URL, response.ErrorResponse{}
}

type CreateCreditCheckoutSessionResponseData struct {
	SessionID  string `json:"session_id" openapi:"example:cs_test_1234567890;"`
	SessionURL string `json:"session_url" openapi:"example:https://checkout.stripe.com/pay/cs_test_1234567890;"`
}
