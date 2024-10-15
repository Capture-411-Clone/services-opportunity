package collector

import (
	"bytes"
	"fmt"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func UnipdfExtractTextFromPDF(pdfData []byte) (string, error) {
	reader := bytes.NewReader(pdfData)
	pdfReader, err := model.NewPdfReader(reader)
	if err != nil {
		return "", fmt.Errorf("error creating PDF reader: %w", err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", fmt.Errorf("error getting number of pages: %w", err)
	}

	hardLimit := 2

	// Hard Limit to 4 pages
	if numPages > hardLimit {
		numPages = hardLimit
	}

	var pagesText []string
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return "", fmt.Errorf("error retrieving page %d: %w", i, err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", fmt.Errorf("error creating extractor for page %d: %w", i, err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			return "", fmt.Errorf("error extracting text from page %d: %w", i, err)
		}

		if text == "" {
			text = "No text found on this page."
		}
		pagesText = append(pagesText, text)
	}

	var text string

	for _, pageText := range pagesText {
		text += pageText + "\n"
	}

	return text, nil
}
