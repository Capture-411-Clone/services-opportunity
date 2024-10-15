package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: account
 * @apiPath: /change-password
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: ChangePasswordRequest
 * @apiResponseRef: ChangePasswordResponse
 * @apiSummary: change password
 * @apiDescription: change password
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) changePassword(ctx echo.Context) error {
	var input = &endpoints.ChangePasswordRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	loginResponse, err := r.service.ChangePassword(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(loginResponse))
}

/*
 * @apiDefine: ChangePasswordResponse
 */
type ChangePasswordResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:201;type:integer;"`
	Data       endpoints.LoginResponseData `json:"data" openapi:"$ref:LoginResponseData;type:string"`
}
