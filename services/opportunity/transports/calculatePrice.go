package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/{id}/price
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: CalculateOpportunityPriceRequestParam
 * @apiResponseRef: CalculateOpportunityPriceResponse
 * @apiSummary: calculate opportunity price
 * @apiDescription: calculate opportunity price
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) calculatePrice(ctx echo.Context) error {
	params := &endpoints.CalculateOpportunityPriceRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	finalPrice, err := r.service.CalculateOpportunityPrice(ctx.Request().Context(), params.ID)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	result := CalculateOpportunityPriceResponseData{
		Price:    finalPrice / 100,
		Currency: "USD",
		Prefix:   "$",
	}

	return ctx.JSON(http.StatusCreated, response.Success(result))
}

/*
 * @apiDefine: CalculateOpportunityPriceResponseData
 */
type CalculateOpportunityPriceResponseData struct {
	Price    float64 `json:"price" openapi:"example:200;type:integer"`
	Currency string  `json:"currency" openapi:"example:USD;type:string"`
	Prefix   string  `json:"prefix" openapi:"example:$;type:string"`
}

/*
 * @apiDefine: CalculateOpportunityPriceResponse
 */
type CalculateOpportunityPriceResponse struct {
	StatusCode int                                   `json:"status_code" openapi:"example:200;type:integer"`
	Data       CheckDuplicateOpportunityResponseData `json:"data" openapi:"$ref:CalculateOpportunityPriceResponseData;type:object"`
}
