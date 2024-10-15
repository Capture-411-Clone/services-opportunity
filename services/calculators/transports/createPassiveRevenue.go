package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/calculators/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: calculators
 * @apiPath: /calculators/passive-revenue
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreatePassiveRevenueCalculatorRequest
 * @apiResponseRef: CreatePassiveRevenueResponse
 * @apiSummary: create passive revenue calculator
 * @apiDescription: create passive revenue calculator
 */
func (r *resource) createPassiveRevenue(ctx echo.Context) error {
	var input = &endpoints.CreatePassiveRevenueCalculatorRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	calculator, err := r.service.CreatePassiveRevenue(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(calculator))
}

/*
 * @apiDefine: CreatePassiveRevenueResponse
 */
type CreatePassiveRevenueResponse struct {
	StatusCode int                   `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.PassiveRevenue `json:"data" openapi:"$ref:PassiveRevenue;type:object"`
}
