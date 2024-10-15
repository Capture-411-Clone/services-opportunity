package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/office/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: office
 * @apiPath: /offices
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateOfficeRequest
 * @apiResponseRef: CreateOfficeResponse
 * @apiSummary: create office
 * @apiDescription: create office
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateOfficeRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	office, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(office))
}

/*
 * @apiDefine: CreateOfficeResponse
 */
type CreateOfficeResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Office `json:"data" openapi:"$ref:Office;type:object"`
}
