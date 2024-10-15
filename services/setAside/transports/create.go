package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/setAside/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: setAside
 * @apiPath: /setAsides
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateSetAsideRequest
 * @apiResponseRef: CreateSetAsideResponse
 * @apiSummary: create setAside
 * @apiDescription: create setAside
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateSetAsideRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	setAside, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(setAside))
}

/*
 * @apiDefine: CreateSetAsideResponse
 */
type CreateSetAsideResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.SetAside `json:"data" openapi:"$ref:SetAside;type:object"`
}
