package welcome

import (
	"fmt"
	"net/http"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Welcome to %s API", config.AppConfig.Name))
	})
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Welcome to %s API V1", config.AppConfig.Name))
	})
}
