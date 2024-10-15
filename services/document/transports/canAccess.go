package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/document/endpoints"
	"github.com/labstack/echo/v4"
)

/*
* @apiTag: document
* @apiPath: /can-access/documents/
* @apiMethod: POST
* @apiStatusCode: 200
* @apiRequestRef: CanAccessDocumentRequest
* @apiResponseRef:CanAccessDocumentResponse
* @apiSummary: check document access
* @apiDescription:check document access
 */
func (r *resource) canAccess(ctx echo.Context) error {
	var input = &endpoints.CanAccessDocumentRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "Please provide valid input data"))
	}

	resStr, err := r.service.CanAccess(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(resStr))
}

/*
 * @apiDefine: CanAccessDocumentResponse
 */
type CanAccessDocumentResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:OK-HasAccess;type:string"`
}
