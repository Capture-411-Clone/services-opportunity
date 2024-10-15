package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: verify
 * @apiPath: /verifications/send
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiRequestRef: SendCodeRequest
 * @apiResponseRef: SendCodeResponse
 * @apiSummary: send verification code
 * @apiDescription: send verification code
 */
func (r resource) sendCode(ctx echo.Context) error {
	var input = &endpoints.SendCodeRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	code, err := r.service.SendCode(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	result := &sendCodeResponseData{
		Code: code,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: SendCodeResponse
 */
type SendCodeResponse struct {
	StatusCode int                  `json:"status_code" openapi:"example:200;type:integer"`
	Data       sendCodeResponseData `json:"data" openapi:"$ref:sendCodeResponseData;type:object;"`
}

/*
 * @apiDefine: sendCodeResponseData
 */
type sendCodeResponseData struct {
	Code string `json:"code" openapi:"example:123456"`
}
