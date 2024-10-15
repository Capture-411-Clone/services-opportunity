package transports

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/opportunity/collector"
	"github.com/labstack/echo/v4"
)

// FIXME: add request body with multipart form data for files
/*
 * @apiTag: opportunity
 * @apiPath: /opportunities/pdf-doc-values
 * @apiMethod: POST
 * @apiStatusCode: 200
 * @apiResponseRef: OpportunityPdfDocValuesResponse
 * @apiSummary: Get Pdf Doc Values
 * @apiDescription: Get Pdf Doc Values
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) pdfDocValues(ctx echo.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	prompt := form.Value["prompt"]

	files := form.File["files"]
	if len(files) == 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "No files found"))
	}

	file := files[0]

	// Open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	pdfData, err := io.ReadAll(src)
	if err != nil {
		r.logger.With(ctx.Request().Context()).Error(err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "Error Opening the file"))
	}

	// Extract text from PDF
	textPages, err := collector.PopplerExtractTextFromPDF(pdfData)

	if err != nil {
		r.logger.With(ctx.Request().Context()).Error(err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "Error extracting text from PDF"))
	}

	limit := 16385 // GPT 3.5-Turbo Token Limit

	shrunkText, withinLimit, tokenCount := collector.ShrinkTextToFitLimit(textPages, limit)
	if withinLimit {
		fmt.Println("******* Original text is within the token limit.")
	} else {
		fmt.Printf("******* Original text exceeds the token limit by %d tokens. Main token count: %d\n", tokenCount-limit, tokenCount)
	}

	// if index 0 of promt does not exists make prompt empty string
	if len(prompt) == 0 {
		prompt = append(prompt, "")
	}

	// Extract values from text
	values, err := collector.GptGetDocValues(shrunkText, prompt[0])
	if err != nil {
		r.logger.With(ctx.Request().Context()).Error(err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(nil, "Error extracting values from text"))
	}

	return ctx.JSON(http.StatusOK, response.Success(values))
}

/*
 * @apiDefine: OpportunityPdfDocValuesResponse
 */
type OpportunityPdfDocValuesResponse struct {
	StatusCode int                       `json:"status_code" openapi:"example:200;type:integer;"`
	Data       collector.GptChoiceValues `json:"data" openapi:"$ref:GptChoiceValues;type:object"`
}
