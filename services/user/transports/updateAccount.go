package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /account/update
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: UpdateUserAccountRequest
 * @apiResponseRef: UpdateUserAccountResponse
 * @apiSummary: update account
 * @apiDescription: update account
 */
func (r resource) updateAccount(ctx echo.Context) error {
	var input = &endpoints.UpdateUserAccountRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "Please check your input values"))
	}

	user, err := r.service.UpdateAccount(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: UpdateUserAccountResponse
 */
type UpdateUserAccountResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.UserData `json:"data" openapi:"$ref:UserData;type:object"`
}
