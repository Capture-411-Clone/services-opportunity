package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/billingportal/session"
)

func (s *service) CreatePortalSession(ctx context.Context) (string, response.ErrorResponse) {
	redirectURL := config.AppConfig.StripeRedirectDomain
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreatePortalSession: User not found, %v", err)
		return "", response.ErrorNotFound("User not found")
	}

	if user.StripeCustomerID == "" {
		s.logger.With(ctx).Errorf("CreatePortalSession: User data is not complete, no StripeCustomerID")
		return "", response.ErrorBadRequest("User data is not complete")
	}

	// Create a Billing Portal Session
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(user.StripeCustomerID),
		ReturnURL: stripe.String(redirectURL), // URL to which the customer will be redirected after they finish using the billing portal.
	}

	sess, err := session.New(params)
	if err != nil {
		s.logger.With(ctx).Errorf("CreatePortalSession: Error creating session, %v", err)
		return "", response.ErrorBadRequest("Error creating session")
	}

	// Return the URL of the session to redirect the user
	return sess.URL, response.ErrorResponse{}
}
