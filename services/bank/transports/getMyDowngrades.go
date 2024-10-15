package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/labstack/echo/v4"
)

func (r resource) getMyDowngrades(ctx echo.Context) error {
	downgrades, err := r.service.GetMyDowngrades(ctx.Request().Context())

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Success(downgrades))
}
