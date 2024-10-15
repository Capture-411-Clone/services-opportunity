package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/encrypt"
)

/*
 * @apiDefine: ResetPasswordRequest
 */
type ResetPasswordRequest struct {
	Email       string `json:"email,omitempty" openapi:"example:JnY6A@example.com;type:string"`
	Password    string `json:"password,omitempty" openapi:"example:1234568;type:string"`
	SessionCode string `json:"session_code,omitempty" openapi:"example:uEmailId;type:string"`
}

func (r *ResetPasswordRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"email":        gocriteria.New("email").Required().Email(),
		"password":     PasswordValidity,
		"session_code": gocriteria.New("session_code").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"email":        "Email",
			"password":     "Password",
			"session_code": "Session Code",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) ResetPassword(ctx context.Context, input ResetPasswordRequest) (LoginResponseData, response.ErrorResponse) {
	var user models.User
	var exists bool
	var res LoginResponseData
	exists, user, err := s.findUser(ctx, input.Email)
	if !exists {
		return res, response.ErrorBadRequest(nil, "No user with this username was found")
	}

	_, responseError := s.checkAndDeleteVerificationBySessionCodeAndEmail(ctx, input.SessionCode, input.Email)
	if responseError.StatusCode != 0 {
		s.logger.With(ctx).Error(err)
		return res, responseError
	}

	var hashedPassword string
	hashedPassword, err = encrypt.HashPassword(input.Password)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.ErrorInternalServerError(err, "An error occurred on the server")
	}

	user.Password = hashedPassword
	user.ShouldChangePassword = false

	err = s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.GormErrorResponse(err, "An error occurred while generating the password")
	}

	accessToken, err := s.generateAccessTokenJWT(user)
	if responseError.StatusCode != 0 {
		s.logger.With(ctx).Error(err)
		return res, response.ErrorBadRequest(err, "An error occurred while generating the token")
	}
	res = LoginResponseData{
		User:        user,
		AccessToken: accessToken,
	}
	return res, response.ErrorResponse{}
}
