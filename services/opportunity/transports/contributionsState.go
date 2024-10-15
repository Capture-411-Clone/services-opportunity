package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/contribution-state
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiResponseRef: ContributionStateResponse
 * @apiSummary: contribution state
 * @apiDescription: contribution state
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) contributionsState(ctx echo.Context) error {
	state, err := r.service.ContributionsState(ctx.Request().Context())
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(state))
}

/*
 * @apiDefine: ContributionsStateResponse
 */
type ContributionsStateResponse struct {
	StatusCode int                                      `json:"status_code" openapi:"example:201;type:integer;"`
	Data       endpoints.ContributionsStateResponseData `json:"data" openapi:"$ref:Opportunity;type:object"`
}
