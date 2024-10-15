package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: verify
 * @apiPath: /verifications/exchange
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: ExchangeRequest
 * @apiResponseRef: ExchangeResponse
 * @apiSummary: exchange verification code
 * @apiDescription: exchange verification code
 */

func (r resource) exchange(ctx echo.Context) error {
	var input = &endpoints.ExchangeRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	sessionCode, err := r.service.Exchange(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	result := &exchangeResponseData{
		SessionCode: sessionCode,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: ExchangeResponse
 */
type ExchangeResponse struct {
	StatusCode int                  `json:"status_code" openapi:"example:201;type:integer;"`
	Data       exchangeResponseData `json:"data" openapi:"$ref:exchangeResponseData;type:object;"`
}

/*
 * @apiDefine: exchangeResponseData
 */
type exchangeResponseData struct {
	SessionCode string `json:"session_code" openapi:"example:123456"`
}
