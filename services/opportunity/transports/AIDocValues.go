package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/collector"
	"github.com/labstack/echo/v4"
)

/*
 * @apiDefine: OpportunityAIDocValuesRequest
 */
type OpportunityAIDocValuesRequest struct {
	FileObjectKey string `json:"file_object_key" openapi:"example:file_object_key;type:string;"`
}

func (c *OpportunityAIDocValuesRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"file_object_key": gocriteria.New("file_object_key").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"file_object_key": "File Object Key",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/ai-doc-values
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiRequestRef: OpportunityAIDocValuesRequest
 * @apiResponseRef: OpportunityAIDocValuesResponse
 * @apiSummary: Get AI Doc Values
 * @apiDescription: Get AI Doc Values
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) AIDocValues(ctx echo.Context) error {
	// Extract values from text
	var input = &OpportunityAIDocValuesRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	values, err := collector.AIGetDocValues(input.FileObjectKey)
	if err != nil {
		r.logger.With(ctx.Request().Context()).Error(err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "Error extracting values from text"))
	}

	return ctx.JSON(http.StatusOK, response.Success(values))
}

/*
 * @apiDefine: OpportunityAIDocValuesResponse
 */
type OpportunityAIDocValuesResponse struct {
	StatusCode int                         `json:"status_code" openapi:"example:200;type:integer;"`
	Data       collector.EntHunterResponse `json:"data" openapi:"$ref:EntHunterResponse;type:object"`
}
