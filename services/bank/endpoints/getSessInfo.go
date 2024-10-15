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
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type GetSessInfoRequestParams struct {
	SessionID string `json:"session_id" openapi:"example:PbRzzvjoOCsfMQ;nullable;in:path"`
}

func (p *GetSessInfoRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"session_id": gocriteria.New("session_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"session_id": "Session ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) GetSessionInfo(ctx context.Context, params GetSessInfoRequestParams) (*stripe.CheckoutSession, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		return &stripe.CheckoutSession{}, response.ErrorNotFound("User not found")
	}

	sess, err := session.Get(params.SessionID, nil)
	if err != nil {
		// Handle the error appropriately.
		s.logger.Error("error getting session info", err)
		return nil, response.ErrorBadRequest(nil, "Session Info not available")
	}

	return sess, response.ErrorResponse{}
}
