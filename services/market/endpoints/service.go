package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters MarketQueryRequestParams) (
		markets []models.Market, err response.ErrorResponse,
	)
	Count(ctx context.Context, params MarketQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateMarketRequest) (market models.Market, err response.ErrorResponse)
	Update(ctx context.Context, id string, input UpdateMarketRequest) (market models.Market, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteMarketsRequest) (response string, err response.ErrorResponse)
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
