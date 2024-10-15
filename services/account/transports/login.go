package transports

import (
	"fmt"
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

/*
* @apiTag: account
* @apiMethod: POST
* @apiStatusCode: 200
* @apiDeprecated: false
* @apiPath: /login
* @apiRequestRef: LoginRequest
* @apiResponseRef: LoginResponse
* @apiSummary: Login User
* @apiDescription: Login User
 */
func (r resource) login(ctx echo.Context) error {
	var input = &endpoints.LoginRequest{}
	errors := input.Validate(ctx.Request())

	fmt.Println(errors)

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}
	loginResponse, err := r.service.Login(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(loginResponse))
}

/*
 * @apiDefine: LoginResponse
 */
type LoginResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:200;type:integer"`
	Data       endpoints.LoginResponseData `json:"data" openapi:"$ref:LoginResponseData;type:object"`
}
