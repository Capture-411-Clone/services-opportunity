package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/agency/endpoints"
	"github.com/labstack/echo/v4"
)

/*
* @apiTag: agency
* @apiPath: /agencies/{id}
* @apiMethod: PUT
* @apiStatusCode: 200
* @apiParametersRef: AgencyRequestUpdateParams
* @apiRequestRef: UpdateAgencyRequest
* @apiResponseRef: UpdateAgencyResponse
* @apiSummary: update agency
* @apiDescription: update agency
* @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.AgencyRequestUpdateParams{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateAgencyRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	agency, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(agency))
}

/*
 * @apiDefine: UpdateAgencyResponse
 */
type UpdateAgencyResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Agency `json:"data" openapi:"$ref:Agency;type:object"`
}
