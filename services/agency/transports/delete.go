package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/agency/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: agency
 * @apiPath: /agencies
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteAgenciesRequest
 * @apiResponseRef: DeleteAgencyResponse
 * @apiSummary: delete agency
 * @apiDescription: delete agency
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteAgenciesRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting agency"))
	}
	deletedCategories, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedCategories))
}

/*
 * @apiDefine: DeleteAgencyResponse
 */
type DeleteAgencyResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	IDs        string `json:"data" openapi:"example:1,2,3;"`
}
