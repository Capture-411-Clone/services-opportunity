package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) createSubscriptionCheckoutSession(ctx echo.Context) error {
	var input = &endpoints.CreateSubscriptionCheckoutSessionRequestBody{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data is invalid"))
	}

	sessionID, sessionURL, err := r.service.CreateSubscriptionCheckoutSession(ctx.Request().Context(), *input)

	res := &endpoints.CreateSubscriptionCheckoutSessionResponseData{
		SessionID:  sessionID,
		SessionURL: sessionURL,
	}

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(res))
}

type CreateSubscriptionCheckoutSessionResponse struct {
	StatusCode int                                                     `json:"statusCode"`
	Data       endpoints.CreateSubscriptionCheckoutSessionResponseData `json:"data" openapi:"$ref:CreateSubscriptionCheckoutSessionResponseData;type:object;"`
}
