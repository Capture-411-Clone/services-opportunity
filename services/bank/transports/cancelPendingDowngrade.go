package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) cancelPendingDowngrade(ctx echo.Context) error {
	params := &endpoints.CancelPendingDowngradeRequestParams{}
	p := map[string]string{
		"downgrade_id": ctx.Param("downgrade_id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	err := r.service.CancelPendingDowngrade(ctx.Request().Context(), *params)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Success("Subscription canceled successfully"))
}
