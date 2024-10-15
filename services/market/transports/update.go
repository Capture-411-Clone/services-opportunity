package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/market/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: market
 * @apiPath: /markets/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateMarketRequestParam
 * @apiRequestRef: UpdateMarketRequest
 * @apiResponseRef: UpdateMarketResponse
 * @apiSummary: update market
 * @apiDescription: update market
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateMarketRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateMarketRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	market, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(market))
}

/*
 * @apiDefine: UpdateMarketResponse
 */
type UpdateMarketResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Market `json:"data" openapi:"$ref:Market;type:object"`
}
