package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters SetAsideQueryRequestParams) (
		setAsides []models.SetAside, err response.ErrorResponse,
	)
	Count(ctx context.Context, params SetAsideQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateSetAsideRequest) (setAside models.SetAside, err response.ErrorResponse)
	Update(ctx context.Context, id string, input UpdateSetAsideRequest) (setAside models.SetAside, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteSetAsideRequest) (response string, err response.ErrorResponse)
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
