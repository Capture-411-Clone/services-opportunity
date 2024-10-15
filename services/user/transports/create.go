package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: user
 * @apiPath: /users
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateUserRequest
 * @apiResponseRef: CreateUserResponse
 * @apiSummary: Create user
 * @apiDescription: Create user
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateUserRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	user, err := r.service.Create(ctx.Request().Context(), *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(user))
}

/*
 * @apiDefine: CreateUserResponse
 */
type CreateUserResponse struct {
	StatusCode int         `json:"status_code"`
	Data       models.User `json:"data" openapi:"$ref:User;type:object"`
}
