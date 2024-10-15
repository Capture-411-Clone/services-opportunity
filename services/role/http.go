package role

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)
	g.GET("/roles/list", res.list)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) list(ctx echo.Context) error {
	c := ctx.Request().Context()
	roles, er := r.service.Query(c)
	if er != nil {
		return er
	}
	return ctx.JSON(http.StatusOK, response.Success(roles))
}
