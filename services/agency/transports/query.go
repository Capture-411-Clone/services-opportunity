package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/agency/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: agency
 * @apiPath: /agencies
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: AgencyQueryRequestParams
 * @apiResponseRef: AgencyQueryResponse
 * @apiSummary: get agency list
 * @apiErrorStatusCodes: 400, 404
 * @apiSecurity: jwtBearerAuth
 */

func (r *resource) query(ctx echo.Context) error {
	var params = &endpoints.AgencyQueryRequestParams{}

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

	agencies, er := r.service.Query(
		ctx.Request().Context(),
		pages.Offset(), pages.Limit(), *params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	pages.Items = agencies

	result := agencyQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      agencies,
	}
	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: agencyQueryResponseData
 */
type agencyQueryResponseData struct {
	Limit      int             `json:"limit" openapi:"example:20"`
	Offset     int             `json:"offset" openapi:"example:0"`
	Page       int             `json:"page" openapi:"example:1"`
	TotalRows  int64           `json:"totalRows" openapi:"example:10"`
	TotalPages int             `json:"totalPages" openapi:"example:1"`
	Items      []models.Agency `json:"items" openapi:"$ref:Agency;type:array;"`
}

/*
 * @apiDefine: AgencyQueryResponse
 */
type AgencyQueryResponse struct {
	StatusCode int                     `json:"status_code" openapi:"example:200;type:integer"`
	Data       agencyQueryResponseData `json:"data" openapi:"$ref:agencyQueryResponseData;type:object"`
}
