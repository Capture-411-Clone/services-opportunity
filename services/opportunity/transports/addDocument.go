package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/{opportunity_id}/documents
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiParametersRef: AddDocumentRequestParam
 * @apiRequestRef: AddDocumentRequest
 * @apiResponseRef: AddDocumentResponse
 * @apiSummary: add document
 * @apiDescription: add document
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) addDocument(ctx echo.Context) error {
	params := &endpoints.AddDocumentRequestParam{}
	p := gocriteria.Params{
		"opportunity_id": ctx.Param("opportunity_id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.AddDocumentRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data not valid"))
	}

	opportunity, err := r.service.AddDocument(ctx.Request().Context(), params.OpportunityID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(opportunity))
}

/*
 * @apiDefine: AddDocumentResponse
 */
type AddDocumentResponse struct {
	StatusCode int                `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Opportunity `json:"data" openapi:"$ref:Opportunity;type:object"`
}
