package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
)

type GetSubscriptionInfoRequestParams struct {
	SubscriptionID string `json:"subscription_id" openapi:"example:PbRzzvjoOCsfMQ;nullable;in:path"`
}

func (p *GetSubscriptionInfoRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"subscription_id": gocriteria.New("subscription_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"subscription_id": "Subscription ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) GetSubscriptionInfo(ctx context.Context, params GetSubscriptionInfoRequestParams) (*stripe.Subscription, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		return &stripe.Subscription{}, response.ErrorNotFound("User not found")
	}

	// Retrieve subscription information from Stripe.
	subscription, err := subscription.Get(params.SubscriptionID, nil)
	if err != nil {
		// Handle the error appropriately.
		s.logger.Error("error getting subscription info", err)
		return nil, response.ErrorBadRequest(nil, "Subscription Info not available")
	}

	return subscription, response.ErrorResponse{}
}
