package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) createCreditCheckoutSession(ctx echo.Context) error {
	var input = &endpoints.CreateCreditCheckoutSessionRequestBody{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data is invalid"))
	}

	sessionID, sessionURL, err := r.service.CreateCreditCheckoutSession(ctx.Request().Context(), *input)

	res := &endpoints.CreateCreditCheckoutSessionResponseData{
		SessionID:  sessionID,
		SessionURL: sessionURL,
	}

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(res))
}

type CreateCreditCheckoutSessionResponse struct {
	StatusCode int                                               `json:"statusCode"`
	Data       endpoints.CreateCreditCheckoutSessionResponseData `json:"data" openapi:"$ref:CreateCreditCheckoutSessionResponseData;type:object;"`
}
