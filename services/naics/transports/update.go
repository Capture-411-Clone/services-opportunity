package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/naics/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: naics
 * @apiPath: /naics/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateNaicsRequestParam
 * @apiRequestRef: UpdateNaicsRequest
 * @apiResponseRef: UpdateNaicsResponse
 * @apiSummary: update naics
 * @apiDescription: update naics
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateNaicsRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateNaicsRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	naics, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(naics))
}

/*
 * @apiDefine: UpdateNaicsResponse
 */
type UpdateNaicsResponse struct {
	StatusCode int          `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Naics `json:"data" openapi:"$ref:Naics;type:object"`
}
