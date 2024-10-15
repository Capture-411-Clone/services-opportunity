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
 * @apiPath: /calculators/capture-cost
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateCaptureCostCalculatorRequest
 * @apiResponseRef: CreateCaptureCostCalculatorResponse
 * @apiSummary: create capture cost calculator
 * @apiDescription: create capture cost calculator
 */
func (r *resource) createCaptureCost(ctx echo.Context) error {
	var input = &endpoints.CreateCaptureCostCalculatorRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	calculator, err := r.service.CreateCaptureCost(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(calculator))
}

/*
 * @apiDefine: CreateCaptureCostCalculatorResponse
 */
type CreateCaptureCostCalculatorResponse struct {
	StatusCode int                `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.CaptureCost `json:"data" openapi:"$ref:CaptureCost;type:object"`
}
