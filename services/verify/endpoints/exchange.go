package endpoints

import (
	"context"
	"net/http"
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: ExchangeRequest
 */
type ExchangeRequest struct {
	Code string `json:"code,omitempty" openapi:"example:123456;type:string"`
}

func (r *ExchangeRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"code": gocriteria.New("code").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"code": "Code",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) Exchange(ctx context.Context, input ExchangeRequest) (string, response.ErrorResponse) {
	var verification models.Verification
	err := s.db.WithContext(ctx).Find(&verification, "code", input.Code).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred while finding the code")
	}
	if verification.Expired() {
		s.logger.Infof("expired: %v", verification.Expired(), verification.ExpiresAt, time.Now())
		return "", response.GormErrorResponse(err, "The code has expired.")
	}

	if verification.Exchanged() {
		return "", response.GormErrorResponse(err, "The code has been used before.")
	}

	verification.ExchangedAt = time.Now()
	if err = s.db.WithContext(ctx).Save(&verification).Error; err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred while saving the code")
	}

	sessionCode := verification.SessionCode
	return sessionCode, response.ErrorResponse{}
}
