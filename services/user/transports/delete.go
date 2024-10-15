package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /users
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteUserRequest
 * @apiResponseRef: DeleteUserResponse
 * @apiSummary: delete user
 * @apiDescription: delete user
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) delete(ctx echo.Context) error {
	input := &endpoints.DeleteUserRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting User"))
	}
	deletedUsers, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedUsers))
}

/*
 * @apiDefine: DeleteUserResponse
 */
type DeleteUserResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
