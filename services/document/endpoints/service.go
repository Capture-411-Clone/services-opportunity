package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Update(ctx context.Context, id string, input UpdateDocumentRequest) (document models.Document, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteDocumentsRequest) (response string, err response.ErrorResponse)
	CanAccess(ctx context.Context, input CanAccessDocumentRequest) (string, response.ErrorResponse)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{db, logger}
}
