package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/order/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: order
 * @apiPath: /orders/{id}/toggle-refund
 * @apiMethod: PATCH
 * @apiStatusCode: 200
 * @apiParametersRef: ToggleRefundRequestParam
 * @apiResponseRef: ToggleRefundResponse
 * @apiSummary: update refund field
 * @apiDescription: update refund field
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) toggleRefund(ctx echo.Context) error {
	params := &endpoints.ToggleRefundRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	user, err := r.service.ToggleRefund(ctx.Request().Context(), params.ID)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: ToggleRefundResponse
 */
type ToggleRefundResponse struct {
	StatusCode int          `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Order `json:"data" openapi:"$ref:Order;type:object"`
}
