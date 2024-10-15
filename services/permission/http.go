package permission

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

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)
	rg.GET("/permissions/access-list", res.accessList)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) accessList(ctx echo.Context) error {
	c := ctx.Request().Context()
	acl, er := r.service.AccessList(c)
	if er != nil {
		return er
	}
	return ctx.JSON(http.StatusOK, response.Success(acl))
}
