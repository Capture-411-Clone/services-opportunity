package endpoints

import (
	"context"
	"errors"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

/*
 * @apiDefine: CheckEmailExistsRequest
 */
type CheckEmailExistsRequest struct {
	Email string `json:"email" openapi:"example:christina@example.com;type:string"`
}

func (r *CheckEmailExistsRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"email": gocriteria.New("email").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"email": "Email",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) CheckEmailExists(ctx context.Context, input CheckEmailExistsRequest) (bool, response.ErrorResponse) {
	var user models.User
	var count int64
	err := s.db.WithContext(ctx).First(&user, "email", input.Email).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.With(ctx).Error(err)
		return false, response.GormErrorResponse(err, "An error occurred while finding the email")
	}
	exists := count > 0
	return exists, response.ErrorResponse{}
}
