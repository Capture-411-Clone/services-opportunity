package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/naics/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: naics
 * @apiPath: /naics
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: NaicsQueryRequestParams
 * @apiResponseRef: NaicsQueryResponse
 * @apiSummary: get naics
 * @apiDescription: get naics
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) query(ctx echo.Context) error {
	var params = &endpoints.NaicsQueryRequestParams{}

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

	naicses, er := r.service.Query(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = naicses

	result := naicsQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      naicses,
	}
	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: naicsQueryResponseData
 */
type naicsQueryResponseData struct {
	Limit      int            `json:"limit" openapi:"example:20"`
	Offset     int            `json:"offset" openapi:"example:0"`
	Page       int            `json:"page" openapi:"example:1"`
	TotalRows  int64          `json:"totalRows" openapi:"example:10"`
	TotalPages int            `json:"totalPages" openapi:"example:1"`
	Items      []models.Naics `json:"items" openapi:"$ref:Naics;type:array;"`
}

/*
 * @apiDefine: NaicsQueryResponse
 */
type NaicsQueryResponse struct {
	StatusCode int                    `json:"status_code" openapi:"example:200;type:integer"`
	Data       naicsQueryResponseData `json:"data" openapi:"$ref:naicsQueryResponseData;type:object"`
}
