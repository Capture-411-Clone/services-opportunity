package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: verify
 * @apiPath: /verifications/phone/check
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiRequestRef: CheckPhoneExistsRequest
 * @apiResponseRef: CheckPhoneExistsResponse
 * @apiSummary: check phone exists
 * @apiDescription: check phone exists
 */
func (r resource) checkPhoneExists(ctx echo.Context) error {
	var input = &endpoints.CheckPhoneExistsRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	exists, err := r.service.CheckPhoneExists(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(exists))
}

/*
 * @apiDefine: CheckPhoneExistsResponse
 */
type CheckPhoneExistsResponse struct {
	StatusCode int  `json:"status_code" openapi:"example:200;type:integer"`
	Data       bool `json:"data" openapi:"example:true"`
}
