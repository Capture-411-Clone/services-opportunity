package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/document/endpoints"
	"github.com/labstack/echo/v4"
)

/*
* @apiTag: document
* @apiPath: /documents/{id}
* @apiMethod: PUT
* @apiStatusCode: 200
* @apiParametersRef: UpdateDocumentRequestParam
* @apiRequestRef: UpdateDocumentRequest
* @apiResponseRef: UpdateDocumentResponse
* @apiSummary: update document
* @apiDescription: update document
* @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateDocumentRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateDocumentRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "the input data is invalid"))
	}

	document, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(document))
}

/*
 * @apiDefine: UpdateDocumentResponse
 */
type UpdateDocumentResponse struct {
	StatusCode int               `json:"status_code" openapi:"example:200;type:integer"`
	Data       []models.Document `json:"data" openapi:"$ref:Document;type:array"`
}
