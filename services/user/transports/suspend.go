package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /users/{id}/suspend/toggle
 * @apiMethod: PATCH
 * @apiStatusCode: 200
 * @apiParametersRef: SuspendRequestParam
 * @apiResponseRef: SuspendResponse
 * @apiSummary: update suspendAt field
 * @apiDescription: update suspendAt field
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) suspend(ctx echo.Context) error {
	params := &endpoints.SuspendRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	user, err := r.service.Suspend(ctx.Request().Context(), params.ID)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: SuspendResponse
 */
type SuspendResponse struct {
	StatusCode int         `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
