package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/order/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/orders", res.query)
	rg.DELETE("/orders", res.delete)
	rg.PATCH("/orders/:id/toggle-refund", res.toggleRefund)
	rg.GET("/orders/opportunity/owned-ids", res.getAllOwnedOpportunityIDs)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
