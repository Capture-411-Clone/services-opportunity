package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/naics/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: naics
 * @apiPath: /naics
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteNaicsRequest
 * @apiResponseRef: DeleteNaicsResponse
 * @apiSummary: delete naics
 * @apiDescription: delete naics
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteNaicsRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting NAICS"))
	}
	deletedNaicses, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedNaicses))
}

/*
 * @apiDefine: DeleteNaicsResponse
 */
type DeleteNaicsResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
