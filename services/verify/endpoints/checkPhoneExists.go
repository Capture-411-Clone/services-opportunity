package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: CheckPhoneExistsRequest
 */
type CheckPhoneExistsRequest struct {
	Phone string `json:"phone,omitempty" openapi:"example:+1-154-8754-25;type:string"`
}

func (r *CheckPhoneExistsRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"phone": gocriteria.New("phone").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"phone": "Phone",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) CheckPhoneExists(ctx context.Context, input CheckPhoneExistsRequest) (bool, response.ErrorResponse) {
	var user models.User
	var count int64
	err := s.db.WithContext(ctx).First(&user, "phone", input.Phone).Count(&count).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return false, response.GormErrorResponse(err, "An error occurred while finding the phone number")
	}
	exists := count > 0
	return exists, response.ErrorResponse{}
}
