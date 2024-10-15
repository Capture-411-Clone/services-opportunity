package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/agency/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: agency
 * @apiPath: /agencies
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateAgencyRequest
 * @apiResponseRef: CreateAgencyResponse
 * @apiSummary: create agency
 * @apiDescription: create agency
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateAgencyRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	agency, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(agency))
}

/*
 * @apiDefine: CreateAgencyResponse
 */
type CreateAgencyResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Agency `json:"data" openapi:"$ref:Agency;type:object"`
}
