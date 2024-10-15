package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/fiscalYear/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: fiscalYear
 * @apiPath: /fiscalYears
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteFiscalYearsRequest
 * @apiResponseRef: DeleteFiscalYearResponse
 * @apiSummary: delete fiscalYear
 * @apiDescription: delete fiscalYear
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteFiscalYearsRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting fiscal year"))
	}
	deletedFiscalYears, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedFiscalYears))
}

/*
 * @apiDefine: DeleteFiscalYearResponse
 */
type DeleteFiscalYearResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
