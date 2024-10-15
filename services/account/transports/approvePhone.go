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
 * @apiPath: /accounts/phone/approve/{code}
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiParametersRef: ApprovePhoneRequestParam
 * @apiResponseRef: ApprovePhoneResponse
 * @apiSummary: approve user phone
 * @apiDescription: approve user phone
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) approvePhone(ctx echo.Context) error {
	params := &endpoints.ApprovePhoneRequestParam{}
	p := gocriteria.Params{
		"code": ctx.Param("code"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	ok, err := r.service.ApprovePhone(ctx.Request().Context(), params.Code)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(ok))
}

/*
 * @apiDefine: ApprovePhoneResponse
 */
type ApprovePhoneResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:201;type:integer;"`
	Data       string `json:"data" openapi:"example:phone approved;type:string"`
}
