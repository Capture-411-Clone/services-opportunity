package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /accounts/{field}/{value}/check-unique
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiParametersRef: GetCheckIsUniqueFieldParams
 * @apiResponseRef: GetCheckIsUniqueFieldResponse
 * @apiSummary: get isUniqueField in user
 * @apiDescription: get isUniqueField in user
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) checkIsUniqueField(ctx echo.Context) error {
	params := &endpoints.GetCheckIsUniqueFieldParams{}
	p := gocriteria.Params{
		"field": ctx.Param("field"),
		"value": ctx.Param("value"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	exists, _, err := r.service.CheckIsUniqueField(ctx.Request().Context(), params.Field, params.Value)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(exists))
}

/*
 * @apiDefine: GetCheckIsUniqueFieldResponse
 */
type GetCheckIsUniqueFieldResponse struct {
	StatusCode int  `json:"status_code" openapi:"example:200;type:integer"`
	Data       bool `json:"data" openapi:"example:true;type:boolean"`
}
