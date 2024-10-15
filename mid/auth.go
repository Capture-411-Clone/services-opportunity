package mid

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/labstack/echo/v4"
)

func CheckInternalAuthToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("Authorization") != config.AppConfig.InternalAuthToken {
				return c.JSON(http.StatusUnauthorized, response.ErrorForbidden(nil, "Internal auth token is not valid."))
			}
			return next(c)
		}
	}
}
