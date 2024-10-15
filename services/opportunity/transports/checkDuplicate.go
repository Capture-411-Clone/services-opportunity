package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/{solicitation_number}/check-duplicate
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: CheckDuplicateOpportunityRequestParam
 * @apiResponseRef: CheckDuplicateOpportunityResponse
 * @apiSummary: check duplidated opportunity
 * @apiDescription: check duplidated opportunity
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) checkDuplicate(ctx echo.Context) error {
	params := &endpoints.CheckDuplicateOpportunityRequestParam{}
	p := gocriteria.Params{
		"solicitation_number": ctx.Param("solicitation_number"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	duplicated, opportunity, err := r.service.CheckDuplicate(ctx.Request().Context(), params.SolicitationNumber)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	result := CheckDuplicateOpportunityResponseData{
		Dulicated:   duplicated,
		Opportunity: opportunity,
	}

	return ctx.JSON(http.StatusCreated, response.Success(result))
}

/*
 * @apiDefine: CheckDuplicateOpportunityResponseData
 */
type CheckDuplicateOpportunityResponseData struct {
	Dulicated   bool               `json:"dulicated" openapi:"example:200;type:integer"`
	Opportunity models.Opportunity `json:"opportunity" openapi:"$ref:Opportunity;type:object"`
}

/*
 * @apiDefine: CheckDuplicateOpportunityResponse
 */
type CheckDuplicateOpportunityResponse struct {
	StatusCode int                                   `json:"status_code" openapi:"example:200;type:integer"`
	Data       CheckDuplicateOpportunityResponseData `json:"data" openapi:"$ref:CheckDuplicateOpportunityResponseData;type:object"`
}
