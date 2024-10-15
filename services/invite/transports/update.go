package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/invite/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: invite
 * @apiPath: /invites/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateInviteRequestParam
 * @apiRequestRef: UpdateInviteRequest
 * @apiResponseRef: UpdateInviteResponse
 * @apiSummary: update invite
 * @apiDescription: update invite
 * @apiSecurity: jwtBearerAuth
 */

func (r resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateInviteRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateInviteRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	invite, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(invite))
}

/*
 * @apiDefine: UpdateInviteResponse
 */
type UpdateInviteResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.Invite `json:"data" openapi:"$ref:Invite;type:object"`
}
