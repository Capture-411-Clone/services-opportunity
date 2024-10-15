package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/naics/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: naics
 * @apiPath: /naics
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateNaicsRequest
 * @apiResponseRef: CreateNaicsResponse
 * @apiSummary: create naics
 * @apiDescription: create naics
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateNaicsRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	naics, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(naics))
}

/*
 * @apiDefine: CreateNaicsResponse
 */
type CreateNaicsResponse struct {
	StatusCode int          `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Naics `json:"data" openapi:"$ref:Naics;type:object"`
}
