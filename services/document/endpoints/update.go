package endpoints

import (
	"context"
	"net/http"
	"strings"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateDocumentRequest
 */
type UpdateDocumentRequest struct {
	FilePath string `json:"file_path" openapi:"example:file.pdf;type:string"`
	Priority int    `json:"priority" openapi:"example:1;type:int"`
}

func (m *UpdateDocumentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"file_path": gocriteria.New("file_path").Required(),
		"priority":  gocriteria.New("priority").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"file_path": "file",
			"priority":  "Priority",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, m)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiDefine: UpdateDocumentRequestParam
 */
type UpdateDocumentRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateDocumentRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id string, input UpdateDocumentRequest) (models.Document, response.ErrorResponse) {
	var documents models.Document
	if !policy.CanUpdateDocument(ctx, documents) {
		s.logger.With(ctx).Error("you do not have permission to edit document a opportunity")
		return models.Document{}, response.ErrorForbidden("you do not have permission to add document a opportunity")
	}
	err := s.db.WithContext(ctx).First(&documents, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Document{}, response.GormErrorResponse(err, "document not find")
	}

	if !strings.Contains(input.FilePath, config.AppConfig.FileSuppressorPrefix) {
		documents.FilePath = input.FilePath
	}

	documents.Priority = input.Priority

	err = s.db.WithContext(ctx).Save(&documents).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Document{}, response.GormErrorResponse(err, "error in document save")
	}

	return documents, response.ErrorResponse{}
}
