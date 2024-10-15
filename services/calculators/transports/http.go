package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/services/calculators/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	g.POST("/calculators/capture-cost", res.createCaptureCost)
	g.POST("/calculators/passive-revenue", res.createPassiveRevenue)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
