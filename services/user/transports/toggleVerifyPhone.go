package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /users/{id}/verifyPhone/toggle
 * @apiMethod: PATCH
 * @apiStatusCode: 200
 * @apiParametersRef: ToggleVerifyPhoneRequestParam
 * @apiResponseRef: ToggleVerifyPhoneResponse
 * @apiSummary: update phoneVerifiedAt field
 * @apiDescription: update phoneVerifiedAt field
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) toggleVerifyPhone(ctx echo.Context) error {
	params := &endpoints.ToggleVerifyPhoneRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	user, err := r.service.ToggleVerifyPhone(ctx.Request().Context(), params.ID)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: ToggleVerifyPhoneResponse
 */
type ToggleVerifyPhoneResponse struct {
	StatusCode int         `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
