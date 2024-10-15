package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/faker"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: ApprovePhoneRequestParam
 */
type ApprovePhoneRequestParam struct {
	Code string `json:"code" openapi:"example:ApproveCode;nullable;in:path"`
}

func (p *ApprovePhoneRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"code": gocriteria.New("code").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"code": "Code",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}
func (s *service) ApprovePhone(ctx context.Context, code string) (string, response.ErrorResponse) {
	var user models.User
	var phone string
	var ok string

	phone, responseError := s.checkAndDeleteVerificationByCode(ctx, code)
	if responseError.StatusCode != 0 {
		s.logger.With(ctx).Error(responseError)
		return ok, responseError
	}

	_, user, err := s.findUser(ctx, phone)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return ok, response.GormErrorResponse(err, "An error occurred in finding the mobile number")
	}

	user.PhoneVerifiedAt = faker.SQLNow()
	err = s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return ok, response.GormErrorResponse(err, "An error occurred in saving the mobile number")
	}

	ok = "phone number verified successfully"

	return ok, response.ErrorResponse{}
}
