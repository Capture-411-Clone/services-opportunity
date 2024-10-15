package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: verify
 * @apiPath: /verifications/email/check
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiRequestRef: CheckEmailExistsRequest
 * @apiResponseRef: CheckEmailExistsResponse
 * @apiSummary: check email exists
 * @apiDescription: check email exists
 */
func (r resource) checkEmailExists(ctx echo.Context) error {
	var input = &endpoints.CheckEmailExistsRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	exists, err := r.service.CheckEmailExists(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(exists))
}

/*
 * @apiDefine: CheckEmailExistsResponse
 */
type CheckEmailExistsResponse struct {
	StatusCode int  `json:"status_code" openapi:"example:201;type:integer;"`
	Data       bool `json:"data" openapi:"example:true"`
}
