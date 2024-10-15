package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: order
 * @apiPath: /orders/opportunity/owned-ids
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiResponseRef: GetAllOwnedOpportunityIDsResponse
 * @apiSummary: get all owned opportunity ids
 * @apiDescription: get all owned opportunity ids
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) getAllOwnedOpportunityIDs(ctx echo.Context) error {

	ids, err := r.service.GetAllOpportunityOwnedIDs(ctx.Request().Context())
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(ids))
}

/*
 * @apiDefine: GetAllOwnedOpportunityIDsResponse
 */
type GetAllOwnedOpportunityIDsResponse struct {
	StatusCode int     `json:"status_code" openapi:"example:200;type:integer"`
	Data       []int64 `json:"data" openapi:"type:array;items:integer"`
}
