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
 * @apiPath: /users/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateUserRequestParam
 * @apiRequestRef: UpdateUserRequest
 * @apiResponseRef: UpdateUserResponse
 * @apiSummary: update user
 * @apiDescription: update user
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateUserRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	var input = &endpoints.UpdateUserRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	user, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(user))
}

/*
 * @apiDefine: UpdateUserResponse
 */
type UpdateUserResponse struct {
	StatusCode int         `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
