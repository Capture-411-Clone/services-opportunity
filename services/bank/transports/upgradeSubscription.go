package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) upgradeSubscription(ctx echo.Context) error {
	var input = &endpoints.UpgradeSubscriptionRequestBody{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data is invalid"))
	}

	subscription, err := r.service.UpgradeSubscription(ctx.Request().Context(), *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Success(subscription))
}
