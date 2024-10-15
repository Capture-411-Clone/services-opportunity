package transports

import (
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)
	g.POST("/verifications/send", res.sendCode)
	g.POST("/verifications/exchange", res.exchange)
	g.POST("/verifications/phone/check", res.checkPhoneExists)
	g.POST("/verifications/email/check", res.checkEmailExists)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
