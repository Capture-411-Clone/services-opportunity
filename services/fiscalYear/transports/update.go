package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/fiscalYear/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: fiscalYear
 * @apiPath: /fiscalYears/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateFiscalYearRequestParam
 * @apiRequestRef: UpdateFiscalYearRequest
 * @apiResponseRef: UpdateFiscalYearResponse
 * @apiSummary: update fiscalYear
 * @apiDescription: update fiscalYear
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateFiscalYearRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateFiscalYearRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	fiscalYear, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(fiscalYear))
}

/*
 * @apiDefine: UpdateFiscalYearResponse
 */
type UpdateFiscalYearResponse struct {
	StatusCode int               `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.FiscalYear `json:"data" openapi:"$ref:FiscalYear;type:object"`
}
