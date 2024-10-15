package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/department/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: department
 * @apiPath: /departments
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: DepartmentQueryRequestParams
 * @apiResponseRef: DepartmentQueryResponse
 * @apiSummary: get department
 * @apiDescription: get department
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) query(ctx echo.Context) error {
	var params = &endpoints.DepartmentQueryRequestParams{}

	errs := params.ValidateQueries(ctx.Request())
	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errs, "invalid input data"))
	}
	if params.OrderBy == "" {
		params.OrderBy = "name"
	}

	if params.Order == "" {
		params.Order = "asc"
	}
	count, err := r.service.Count(ctx.Request().Context(), *params)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages := pagination.NewFromRequest(ctx.Request(), int(count))

	departments, er := r.service.Query(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = departments

	result := departmentQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      departments,
	}
	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: departmentQueryResponseData
 */
type departmentQueryResponseData struct {
	Limit      int                 `json:"limit" openapi:"example:20"`
	Offset     int                 `json:"offset" openapi:"example:0"`
	Page       int                 `json:"page" openapi:"example:1"`
	TotalRows  int64               `json:"totalRows" openapi:"example:10"`
	TotalPages int                 `json:"totalPages" openapi:"example:1"`
	Items      []models.Department `json:"items" openapi:"$ref:Department;type:array;"`
}

/*
 * @apiDefine: DepartmentQueryResponse
 */
type DepartmentQueryResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:200;type:integer"`
	Data       departmentQueryResponseData `json:"data" openapi:"$ref:departmentQueryResponseData;type:object"`
}
