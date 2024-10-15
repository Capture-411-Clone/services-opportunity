package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)
	g.POST("/login", res.login)
	g.POST("/register", res.register)
	g.POST("/resetPassword", res.resetPassword)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/userinfo", res.userInfo)
	rg.PATCH("/change-password", res.changePassword)

	rg.POST("/accounts/approve-phone/:code", res.approvePhone)
	rg.POST("/accounts/approve-email/:token", res.approveEmail)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
