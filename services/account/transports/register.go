package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: account
 * @apiPath: /register
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: RegisterRequest
 * @apiResponseRef: RegisterResponse
 * @apiSummary: register
 * @apiDescription: register
 */
func (r resource) register(ctx echo.Context) error {
	var input = &endpoints.RegisterRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	resp, err := r.service.Register(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(resp))
}

/*
 * @apiDefine: RegisterResponse
 */
type RegisterResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:201;type:integer;"`
	Data       endpoints.LoginResponseData `json:"data" openapi:"$ref:LoginResponseData;type:object"`
}
