package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters FiscalYearQueryRequestParams) (
		fiscalYears []models.FiscalYear, err response.ErrorResponse,
	)
	Count(ctx context.Context, params FiscalYearQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateFiscalYearRequest) (fiscalYear models.FiscalYear, err response.ErrorResponse)
	Update(ctx context.Context, id string, input UpdateFiscalYearRequest) (fiscalYear models.FiscalYear, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteFiscalYearsRequest) (response string, err response.ErrorResponse)
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
