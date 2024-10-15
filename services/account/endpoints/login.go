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
 * @apiDefine: LoginResponseData
 */
type LoginResponseData struct {
	User        models.User `json:"user" openapi:"$ref:User"`
	AccessToken string      `json:"access_token" openapi:"example:eyJhbcCI6IkpXVCJ9.eyWI2ZjYtZjYwZjYwZjYwZjYwIiwiaWF0IjoxNjI5MjUwNjQyLCJle"`
}

/*
 * @apiDefine: LoginRequest
 */
type LoginRequest struct {
	Username string `json:"username,omitempty" openapi:"example:christina.mee@gmail.com"`
	Password string `json:"password,omitempty" openapi:"example:password"`
}

func (r *LoginRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"username": gocriteria.New("username").MinLength(3).MaxLength(128).Required(),
		"password": gocriteria.New("password").MinLength(8).MaxLength(128).Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"username": "Username",
			"password": "Password",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)
	_, user, err := s.findUser(ctx, username)
	if err != nil {
		logger.Error("authenticate the user has problem in findUser() method", err)
	}
	if user.ValidatePassword(password) {
		logger.Infof("authentication successful")
		return user
	}
	logger.Infof("authentication failed")

	return nil
}

func (s *service) Login(ctx context.Context, input LoginRequest) (LoginResponseData, response.ErrorResponse) {
	identity := s.authenticate(ctx, input.Username, input.Password)
	if identity != nil {
		token, errr := s.generateAccessTokenJWT(identity)
		if errr != nil {
			s.logger.With(ctx).Error(errr)
			return LoginResponseData{}, response.ErrorInternalServerError(errr, "An error occurred while generating the token")
		}

		var user models.User
		_, user, err := s.findUser(ctx, input.Username)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return LoginResponseData{}, response.GormErrorResponse(err, "Incorrect username or password")
		}

		if user.SuspendedAt.Valid {
			s.logger.With(ctx).Error(err)
			return LoginResponseData{}, response.ErrorUnAuthorized(nil, "your account has been suspended")
		}

		loginResponse := LoginResponseData{
			User:        user,
			AccessToken: token,
		}
		return loginResponse, response.ErrorResponse{}
	}
	return LoginResponseData{}, response.ErrorBadRequest(nil, "Incorrect username or password")
}
