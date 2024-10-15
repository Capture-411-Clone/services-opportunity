package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/staff-report
 * @apiMethod: GET
 * @apiStatusCode: 200
  * @apiParametersRef: OpportunityQueryRequestParams
 * @apiResponseRef: StaffReportResponse
 * @apiSummary: staff report
 * @apiDescription: staff report
 * @apiSecurity: jwtBearerAuth
*/
func (r *resource) staffReport(ctx echo.Context) error {
	var params = &endpoints.StaffReportQueryRequestParams{}

	errs := params.ValidateQueries(ctx.Request())
	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errs, "invalid input data"))
	}
	if params.OrderBy == "" {
		params.OrderBy = "id"
	}

	if params.Order == "" {
		params.Order = "desc"
	}

	count, err := r.service.StaffReportCount(ctx.Request().Context(), *params)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages := pagination.NewFromRequest(ctx.Request(), int(count))

	reports, er := r.service.StaffReport(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = reports

	result := staffReportResponseBodyData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      reports,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: staffReportResponseBodyData
 */
type staffReportResponseBodyData struct {
	Limit      int                                 `json:"limit" openapi:"example:20"`
	Offset     int                                 `json:"offset" openapi:"example:0"`
	Page       int                                 `json:"page" openapi:"example:1"`
	TotalRows  int64                               `json:"totalRows" openapi:"example:10"`
	TotalPages int                                 `json:"totalPages" openapi:"example:1"`
	Items      []endpoints.StaffReportResponseData `json:"items" openapi:"$ref:StaffReportData;type:array;"`
}

/*
 * @apiDefine: StaffReportResponse
 */
type StaffReportResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:200;type:integer"`
	Data       staffReportResponseBodyData `json:"data" openapi:"$ref:staffReportResponseBodyData;type:object"`
}
