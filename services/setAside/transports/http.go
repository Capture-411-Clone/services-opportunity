package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/setAside/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/setAsides", res.query)
	rg.POST("/setAsides", res.create)
	rg.Match([]string{"PUT", "PATCH"}, "/setAsides/:id", res.update)
	rg.DELETE("/setAsides", res.delete)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
