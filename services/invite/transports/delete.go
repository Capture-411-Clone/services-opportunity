package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/invite/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: invite
 * @apiPath: /invites
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteInviteRequest
 * @apiResponseRef: DeleteInviteResponse
 * @apiSummary: delete invite
 * @apiDescription: delete invite
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteInviteRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting Invite"))
	}
	deletedInvites, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	result := deleteInviteResponseData{
		IDs: deletedInvites,
	}
	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: DeleteInviteResponse
 */
type DeleteInviteResponse struct {
	StatusCode int                      `json:"status_code" openapi:"example:200;type:integer"`
	Data       deleteInviteResponseData `json:"data" openapi:"$ref:deleteInviteResponseData;type:object;"`
}

/*
 * @apiDefine: deleteInviteResponseData
 */
type deleteInviteResponseData struct {
	IDs []int `json:"ids" openapi:"example:[1,2,3];type:array;"`
}
