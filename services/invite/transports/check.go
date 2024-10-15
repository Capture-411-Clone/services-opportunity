package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: invite
 * @apiPath: /invites
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: CheckInviteResponseParams
 * @apiResponseRef: CheckInviteResponse
 * @apiSummary: check invite
 * @apiDescription: check invite
 */
func (r resource) check(ctx echo.Context) error {
	code := ctx.Param("code")
	result, err := r.service.Check(ctx.Request().Context(), code)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: CheckInviteResponse
 */
type CheckInviteResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Invite `json:"data" openapi:"$ref:Invite;type:object;"`
}
