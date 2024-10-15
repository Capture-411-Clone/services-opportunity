package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /users/approvePolicy
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiResponseRef: ApprovePolicyResponse
 * @apiSummary: approve policy
 * @apiDescription: approve policy
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) approvePolicy(ctx echo.Context) error {

	user, err := r.service.ApprovePolicy(ctx.Request().Context())

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: ApprovePolicyResponse
 */
type ApprovePolicyResponse struct {
	StatusCode int         `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
