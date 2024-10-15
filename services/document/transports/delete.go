package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/document/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: document
 * @apiPath: /documents
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteDocumentsRequest
 * @apiResponseRef: DeleteDocumentsResponse
 * @apiSummary: delete document
 * @apiDescription: delete document
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteDocumentsRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting document"))
	}
	deletedMedia, err := r.service.Delete(ctx.Request().Context(), *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedMedia))
}

/*
 * @apiDefine: DeleteDocumentsResponse
 */
type DeleteDocumentsResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,2,3;"`
}
