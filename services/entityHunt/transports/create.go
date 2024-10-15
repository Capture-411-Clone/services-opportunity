package transports

import (
	"fmt"
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/entityHunt/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: entityHunt
 * @apiPath: /entity-hunts
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateEntityHuntRequest
 * @apiResponseRef: CreateEntityHuntResponse
 * @apiSummary: create entityHunt
 * @apiDescription: create entityHunt
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateEntityHuntRequest{}

	errors := input.Validate(ctx.Request())
	fmt.Println("errrrrrrrrrrrrrrrrrrrr", errors)
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	entityHunt, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(entityHunt))
}

/*
 * @apiDefine: CreateEntityHuntResponse
 */
type CreateEntityHuntResponse struct {
	StatusCode int               `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.EntityHunt `json:"data" openapi:"$ref:EntityHunt;type:object;"`
}
