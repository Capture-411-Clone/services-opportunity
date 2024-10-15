package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: account
 * @apiPath: /resetPassword
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: ResetPasswordRequest
 * @apiResponseRef: ResetPasswordResponse
 * @apiSummary: reset password
 * @apiDescription: reset password
 */
func (r resource) resetPassword(ctx echo.Context) error {
	var input = &endpoints.ResetPasswordRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	loginResponse, err := r.service.ResetPassword(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(loginResponse))
}

/*
 * @apiDefine: ResetPasswordResponse
 */
type ResetPasswordResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:201;type:integer;"`
	Data       endpoints.LoginResponseData `json:"data" openapi:"$ref:LoginResponseData;type:object"`
}
