package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/office/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: office
 * @apiPath: /offices
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: OfficeQueryRequestParams
 * @apiResponseRef: OfficeQueryResponse
 * @apiSummary: get office
 * @apiDescription: get office
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) query(ctx echo.Context) error {
	var params = &endpoints.OfficeQueryRequestParams{}

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

	offices, er := r.service.Query(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = offices

	result := officeQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      offices,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: officeQueryResponseData
 */
type officeQueryResponseData struct {
	Limit      int             `json:"limit" openapi:"example:20"`
	Offset     int             `json:"offset" openapi:"example:0"`
	Page       int             `json:"page" openapi:"example:1"`
	TotalRows  int64           `json:"totalRows" openapi:"example:10"`
	TotalPages int             `json:"totalPages" openapi:"example:1"`
	Items      []models.Office `json:"items" openapi:"$ref:Office;type:array;"`
}

/*
 * @apiDefine: OfficeQueryResponse
 */
type OfficeQueryResponse struct {
	StatusCode int                     `json:"status_code" openapi:"example:200;type:integer"`
	Data       officeQueryResponseData `json:"data" openapi:"$ref:officeQueryResponseData;type:object"`
}
