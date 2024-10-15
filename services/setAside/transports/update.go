package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/setAside/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: setAside
 * @apiPath: /setAsides/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateSetAsideRequestParam
 * @apiRequestRef: UpdateSetAsideRequest
 * @apiResponseRef: UpdateSetAsideResponse
 * @apiSummary: update setAside
 * @apiDescription: update setAside
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateSetAsideRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateSetAsideRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	setAside, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(setAside))
}

/*
 * @apiDefine: UpdateSetAsideResponse
 */
type UpdateSetAsideResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.SetAside `json:"data" openapi:"$ref:SetAside;type:object"`
}
