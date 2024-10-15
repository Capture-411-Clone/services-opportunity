package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: account
 * @apiPath: /userinfo
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiResponseRef: UserInfoResponse
 * @apiSummary: get user info
 * @apiDescription: get user info
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) userInfo(ctx echo.Context) error {
	accessToken := ctx.Request().Header.Get("Authorization")
	accessToken = accessToken[len("Bearer "):]

	user, err := r.service.UserInfo(ctx.Request().Context(), accessToken)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(user))
}

/*
 * @apiDefine: UserInfoResponse
 */
type UserInfoResponse struct {
	StatusCode int         `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
