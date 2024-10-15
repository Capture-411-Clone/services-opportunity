package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	CreateCaptureCost(ctx context.Context, input CreateCaptureCostCalculatorRequest) (calculator models.CaptureCost, err response.ErrorResponse)
	CreatePassiveRevenue(ctx context.Context, input CreatePassiveRevenueCalculatorRequest) (models.PassiveRevenue, response.ErrorResponse)
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
