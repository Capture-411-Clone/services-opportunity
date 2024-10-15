package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/fiscalYear/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: fiscalYear
 * @apiPath: /fiscalYears
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateFiscalYearRequest
 * @apiResponseRef: CreateFiscalYearResponse
 * @apiSummary: create fiscalYear
 * @apiDescription: create fiscalYear
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateFiscalYearRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	fiscalYear, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(fiscalYear))
}

/*
 * @apiDefine: CreateFiscalYearResponse
 */
type CreateFiscalYearResponse struct {
	StatusCode int               `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.FiscalYear `json:"data" openapi:"$ref:FiscalYear;type:object"`
}
