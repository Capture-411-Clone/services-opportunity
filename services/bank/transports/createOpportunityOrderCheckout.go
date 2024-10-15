package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) createOpportunityOrderCheckoutSession(ctx echo.Context) error {
	var input = &endpoints.CreateOpportunityOrderCheckoutSessionRequestBody{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data is invalid"))
	}

	sessionID, sessionURL, err := r.service.CreateOpportunityOrderCheckoutSession(ctx.Request().Context(), *input)

	res := &endpoints.CreateOpportunityOrderCheckoutSessionResponseData{
		SessionID:  sessionID,
		SessionURL: sessionURL,
	}

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(res))
}

type CreateOpportunityOrderCheckoutSessionResponse struct {
	StatusCode int                                                         `json:"statusCode"`
	Data       endpoints.CreateOpportunityOrderCheckoutSessionResponseData `json:"data" openapi:"$ref:CreateOpportunityOrderCheckoutSessionResponseData;type:object;"`
}
