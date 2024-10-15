package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: account
 * @apiPath: /accounts/approve-email/{token}
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiParametersRef: ApproveEmailRequestParam
 * @apiResponseRef: ApproveEmailResponse
 * @apiSummary: approve user email
 * @apiDescription: approve user email
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) approveEmail(ctx echo.Context) error {
	params := &endpoints.ApproveEmailRequestParam{}
	p := gocriteria.Params{
		"token": ctx.Param("token"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	ok, err := r.service.ApprovePhone(ctx.Request().Context(), params.Token)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(ok))
}

/*
 * @apiDefine: ApproveEmailResponse
 */
type ApproveEmailResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:201;type:integer;"`
	Data       string `json:"data" openapi:"example:email approved;type:string"`
}
