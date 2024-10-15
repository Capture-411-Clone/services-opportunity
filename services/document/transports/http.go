package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/document/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	// ------------Internal---------------
	ig := r.Group(prefix)
	ig.Use(mid.CheckInternalAuthToken())
	ig.POST("/internal/documents/can-assess/document", res.canAccess)

	// -------------Public----------------
	rg := r.Group(prefix)
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.Match([]string{"PUT", "PATCH"}, "/documents/:id", res.update)
	rg.DELETE("/documents", res.delete)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
