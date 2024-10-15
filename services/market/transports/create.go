package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/market/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: market
 * @apiPath: /markets
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateMarketRequest
 * @apiResponseRef: CreateMarketResponse
 * @apiSummary: create market
 * @apiDescription: create market
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateMarketRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	market, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(market))
}

/*
 * @apiDefine: CreateMarketResponse
 */
type CreateMarketResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Market `json:"data" openapi:"$ref:Market;type:object"`
}
