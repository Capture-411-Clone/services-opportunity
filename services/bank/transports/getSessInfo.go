package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) getSessInfo(ctx echo.Context) error {
	params := &endpoints.GetSessInfoRequestParams{}
	p := map[string]string{
		"session_id": ctx.Param("session_id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	prices, err := r.service.GetSessionInfo(ctx.Request().Context(), *params)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Success(prices))
}
