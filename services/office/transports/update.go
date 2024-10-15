package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/office/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: office
 * @apiPath: /offices/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateOfficeRequestParam
 * @apiRequestRef: UpdateOfficeRequest
 * @apiResponseRef: UpdateOfficeResponse
 * @apiSummary: update office
 * @apiDescription: update office
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateOfficeRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateOfficeRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	office, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(office))
}

/*
 * @apiDefine: UpdateOfficeResponse
 */
type UpdateOfficeResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Office `json:"data" openapi:"$ref:Office;type:object"`
}
