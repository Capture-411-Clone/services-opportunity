package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters OrderQueryRequestParams) (
		orders []models.Order, err response.ErrorResponse,
	)
	Count(ctx context.Context, params OrderQueryRequestParams) (count int64, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteOrderRequest) (response string, err response.ErrorResponse)

	GetAllOpportunityOwnedIDs(ctx context.Context) ([]int64, response.ErrorResponse)

	ToggleRefund(ctx context.Context, orderID string) (models.Order, response.ErrorResponse)
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
