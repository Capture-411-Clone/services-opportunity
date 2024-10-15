package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteOpportunityRequest
 * @apiResponseRef: DeleteOpportunityResponse
 * @apiSummary: delete opportunity
 * @apiDescription: delete opportunity
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteOpportunityRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting opportunity"))
	}
	deletedOpportunities, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedOpportunities))
}

/*
 * @apiDefine: DeleteOpportunityResponse
 */
type DeleteOpportunityResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
