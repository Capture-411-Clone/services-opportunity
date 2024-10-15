package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateOpportunityRequest
 * @apiResponseRef: CreateOpportunityResponse
 * @apiSummary: create opportunity
 * @apiDescription: create opportunity
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateOpportunityRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	opportunity, err := r.service.Create(ctx.Request().Context(), *input, false)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(opportunity))
}

/*
 * @apiDefine: CreateOpportunityResponse
 */
type CreateOpportunityResponse struct {
	StatusCode int                `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Opportunity `json:"data" openapi:"$ref:Opportunity;type:object"`
}
