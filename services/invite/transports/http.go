package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/invite/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)
	g.GET("/invites/:code/check", res.check)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/invites", res.query)
	rg.POST("/invites", res.create)
	rg.Match([]string{"PUT", "PATCH"}, "/invites/:id", res.update)
	rg.DELETE("/invites", res.delete)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
