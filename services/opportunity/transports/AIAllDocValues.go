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
type OpportunityAIAllDocValuesRequest struct {
	FileObjectKey string `json:"file_object_key" openapi:"example:file_object_key;type:string;"`
}

func (c *OpportunityAIAllDocValuesRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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
 * @apiPath: /opportunities/ai-all-doc-values
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiRequestRef: OpportunityAIAllDocValuesRequest
 * @apiResponseRef: OpportunityAIAllDocValuesResponse
 * @apiSummary: Get AI All Doc Values
 * @apiDescription: Get AI All Doc Values
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) AIAllDocValues(ctx echo.Context) error {
	// Extract values from text
	var input = &OpportunityAIAllDocValuesRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	values, err := collector.AIGetAllDocValues(input.FileObjectKey)
	if err != nil {
		r.logger.With(ctx.Request().Context()).Error(err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "Error extracting values from text"))
	}

	return ctx.JSON(http.StatusOK, response.Success(values))
}

/*
 * @apiDefine: OpportunityAIAllDocValuesResponse
 */
type OpportunityAIAllDocValuesResponse struct {
	StatusCode int                            `json:"status_code" openapi:"example:200;type:integer;"`
	Data       collector.EntHunterAllResponse `json:"data" openapi:"$ref:EntHunterAllResponse;type:object"`
}
