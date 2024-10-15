package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: OpportunityQueryRequestParams
 * @apiResponseRef: OpportunityQueryResponse
 * @apiSummary: get opportunity
 * @apiDescription: get opportunity
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) query(ctx echo.Context) error {
	var params = &endpoints.OpportunityQueryRequestParams{}

	errs := params.ValidateQueries(ctx.Request())
	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errs, "invalid input data"))
	}
	if params.OrderBy == "" {
		params.OrderBy = "title"
	}

	if params.Order == "" {
		params.Order = "asc"
	}
	count, err := r.service.Count(ctx.Request().Context(), *params)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages := pagination.NewFromRequest(ctx.Request(), int(count))

	opportunities, er := r.service.Query(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = opportunities

	result := opportunityQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      opportunities,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: opportunityQueryResponseData
 */
type opportunityQueryResponseData struct {
	Limit      int                  `json:"limit" openapi:"example:20"`
	Offset     int                  `json:"offset" openapi:"example:0"`
	Page       int                  `json:"page" openapi:"example:1"`
	TotalRows  int64                `json:"totalRows" openapi:"example:10"`
	TotalPages int                  `json:"totalPages" openapi:"example:1"`
	Items      []models.Opportunity `json:"items" openapi:"$ref:Opportunity;type:array;"`
}

/*
 * @apiDefine: OpportunityQueryResponse
 */
type OpportunityQueryResponse struct {
	StatusCode int                          `json:"status_code" openapi:"example:200;type:integer"`
	Data       opportunityQueryResponseData `json:"data" openapi:"$ref:opportunityQueryResponseData;type:object"`
}
