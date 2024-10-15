package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/encrypt"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: ChangePasswordRequest
 */
type ChangePasswordRequest struct {
	Password    string `json:"password,omitempty" openapi:"example:1234568;type:string"`
	NewPassword string `json:"new_password,omitempty" openapi:"example:1234568;type:string"`
}

func (r *ChangePasswordRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"password":     gocriteria.New("password").Required(),
		"new_password": PasswordValidity,
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"password":     "Password",
			"new_password": "New Password",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) ChangePassword(
	ctx context.Context, input ChangePasswordRequest,
) (
	LoginResponseData,
	response.ErrorResponse,
) {
	var user models.User

	var loginResponse LoginResponseData

	Id := policy.ExtractIdClaim(ctx)
	err := s.db.WithContext(ctx).First(&user, "id", Id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return loginResponse, response.GormErrorResponse(err, "An error occurred in the database")
	}

	if !user.ValidatePassword(input.Password) {
		s.logger.With(ctx).Error("Incorrect password")
		return LoginResponseData{}, response.ErrorBadRequest(err, "Incorrect password")
	}

	var hashedPassword string
	hashedPassword, err = encrypt.HashPassword(input.NewPassword)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return loginResponse, response.ErrorInternalServerError(err, "An error occurred on the server.")
	}

	user.Password = hashedPassword
	user.ShouldChangePassword = false

	err = s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		err := response.GormErrorResponse(err, "An error occurred while saving.")
		return loginResponse, err
	}

	accessTkn, err := s.generateAccessTokenJWT(user)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return loginResponse, response.ErrorInternalServerError(err, "An error occurred while generating the token")
	}

	loginResponse = LoginResponseData{
		User:        user,
		AccessToken: accessTkn,
	}
	return loginResponse, response.ErrorResponse{}
}
