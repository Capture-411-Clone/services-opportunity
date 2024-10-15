package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/mass-import
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: ImportOpportunitiesRequest
 * @apiResponseRef: ImportOpportunityResponse
 * @apiSummary: import opportunities
 * @apiDescription: import opportunities
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) importerOpportunities(ctx echo.Context) error {
	var input = &endpoints.ImportOpportunitiesRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	err := r.service.ImportOpportunities(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "successful",
	})
}

/*
 * @apiDefine: ImportOpportunityResponse
 */
type ImportOpportunityResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:201;type:integer;"`
	Data       string `json:"data" openapi:"example:successful;type:string"`
}
