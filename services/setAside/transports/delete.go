package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/setAside/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: setAside
 * @apiPath: /setAsides
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteSetAsideRequest
 * @apiResponseRef: DeleteSetAsideResponse
 * @apiSummary: delete setAside
 * @apiDescription: delete setAside
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteSetAsideRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting office"))
	}
	deletedSetAsides, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedSetAsides))
}

/*
 * @apiDefine: DeleteSetAsideResponse
 */
type DeleteSetAsideResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
