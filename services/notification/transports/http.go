package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/notification/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.POST("/notifications", res.create)
	rg.PATCH("/notifications/:id", res.update)
	rg.DELETE("/notifications", res.delete)
	rg.GET("/notifications", res.query)

}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
