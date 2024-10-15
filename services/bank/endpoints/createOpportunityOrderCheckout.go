package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/calculator"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
)

type CreateOpportunityOrderCheckoutSessionRequestBody struct {
	OpportunityID int64  `json:"opportunity_id"`
	RedirectPath  string `json:"redirect_path"`
}

func (data *CreateOpportunityOrderCheckoutSessionRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"opportunity_id": gocriteria.New("opportunity_id").Required(),
		"redirect_path":  gocriteria.New("redirect_path").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"opportunity_id": "Opportunity ID",
			"redirect_path":  "Redirect Path",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) CreateOpportunityOrderCheckoutSession(ctx context.Context, input CreateOpportunityOrderCheckoutSessionRequestBody) (string, string, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process: User not found")
	}

	// Get opportunity
	var opportunity models.Opportunity
	err = s.db.Model(&models.Opportunity{}).Where("id = ?", input.OpportunityID).First(&opportunity).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process: Opportunity not found")
	}

	var prevOrders []models.Order
	err = s.db.Model(models.Order{}).Where("user_id = ? AND opportunity_id = ? AND paid_at IS NOT NULL AND refunded_at IS NULL", theID, input.OpportunityID).Find(&prevOrders).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Checking Prev Orders Failed: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process: Checking Prev Orders Failed")
	}

	if len(prevOrders) > 0 {
		s.logger.With(ctx).Error("CreateOpportunityOrderCheckout: User already owns this opportunity")
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process: You already own this opportunity")
	}

	// if opportunity.CrawlerContractValue == "" && opportunity.UserContractValue == "" {
	// 	s.logger.With(ctx).Error("CreateOpportunityOrderCheckout: CrawlerContractValue and UserContractValue is empty")
	// 	return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process: Opportunity contract value is not available")
	// }

	// convert opportunity crawlerContractValue string to float64
	contractValueFloat64 := 0.0
	if opportunity.CrawlerContractValue != "" {
		contractValueFloat64, err = strconv.ParseFloat(opportunity.CrawlerContractValue, 64)
		if err != nil {
			s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
			return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
		}
	}

	userContractValueFloat64 := 0.0
	if opportunity.UserContractValue != "" {
		userContractValueFloat64, err = strconv.ParseFloat(opportunity.UserContractValue, 64)
		if err != nil {
			s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
			return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
		}
	}

	// Decide which value to use
	var finalValue float64
	if userContractValueFloat64 != 0.0 {
		finalValue = userContractValueFloat64
	} else {
		finalValue = contractValueFloat64
	}

	finalPrice := calculator.CalculateOpportunityPrice(finalValue)

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

	fmt.Println("input.RedirectPath: ", input.RedirectPath)

	successRedirectPath, err := addQueryParameters(input.RedirectPath, map[string]string{
		"payment-status": "success",
		// "session-id":     "{CHECKOUT_SESSION_ID}",
	})
	if err != nil {
		s.logger.With(ctx).Errorf("addQueryParameters: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	cancelRedirectPath, err := addQueryParameters(input.RedirectPath, map[string]string{
		"payment-status": "fail",
		"canceled":       "true",
	})
	if err != nil {
		s.logger.With(ctx).Errorf("addQueryParameters: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}

	fmt.Println("successRedirectPath: ", successRedirectPath)

	fmt.Println("cancelRedirectPath: ", cancelRedirectPath)

	// create order in database to put the order id somewhere to get it back in webhook
	order := models.Order{
		PriceAmount:   finalPrice,
		OpportunityID: input.OpportunityID,
		UserID:        theID,
	}

	if err := s.db.Save(&order).Error; err != nil {
		s.logger.With(ctx).Errorf("Save Order Error: %v", err)
		return "", "", response.ErrorBadRequest(nil, "Could not create order")
	}

	// Adjust checkoutParams for one-time purchase
	checkoutParams := &stripe.CheckoutSessionParams{
		Metadata: map[string]string{
			"order_id": strconv.Itoa(int(order.ID)),
		},
		PaymentMethodTypes: stripe.StringSlice([]string{string(stripe.PaymentMethodTypeCard)}),
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		Customer:           stripe.String(customerID),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(opportunity.Title),
					},
					UnitAmountDecimal: stripe.Float64(finalPrice),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(successRedirectPath),
		CancelURL:  stripe.String(cancelRedirectPath),
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

type CreateOpportunityOrderCheckoutSessionResponseData struct {
	SessionID  string `json:"session_id" openapi:"example:cs_test_1234567890;"`
	SessionURL string `json:"session_url" openapi:"example:https://checkout.stripe.com/pay/cs_test_1234567890;"`
}

func addQueryParameters(path string, params map[string]string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
