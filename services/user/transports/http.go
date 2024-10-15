package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/users", res.query)
	rg.POST("/users", res.create)
	rg.Match([]string{"PUT", "PATCH"}, "/users/:id", res.update)
	rg.DELETE("/users", res.delete)

	rg.PATCH("/users/:id/suspend/toggle", res.suspend)
	rg.PATCH("/users/:id/verifyEmail/toggle", res.toggleVerifyEmail)
	rg.PATCH("/users/:id/verifyPhone/toggle", res.toggleVerifyPhone)

	rg.POST("/users/approvePolicy", res.approvePolicy)

	rg.POST("/accounts/update", res.updateAccount)

	rg.GET("/accounts/:field/:value/check-unique", res.checkIsUniqueField)

	rg.GET("/opportunities/staff-report", res.staffReport)
}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
