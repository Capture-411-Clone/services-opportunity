package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/invite/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: invite
 * @apiPath: /invites
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateInviteRequest
 * @apiResponseRef: CreateInviteResponse
 * @apiSummary: create invite
 * @apiDescription: create invite
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateInviteRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	invite, err := r.service.Create(ctx.Request().Context(), *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Created(invite))
}

/*
 * @apiDefine: CreateInviteResponse
 */
type CreateInviteResponse struct {
	StatusCode int           `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Invite `json:"data" openapi:"$ref:Invite;type:object;"`
}
