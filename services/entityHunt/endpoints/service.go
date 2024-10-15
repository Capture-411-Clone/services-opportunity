package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters EntityHuntQueryRequestParams) (
		entityHunts []models.EntityHunt, err response.ErrorResponse,
	)
	Count(ctx context.Context, params EntityHuntQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateEntityHuntRequest) (entityHunt models.EntityHunt, err response.ErrorResponse)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{
		db:     db,
		logger: logger,
	}
}
