package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	AddDocument(ctx context.Context, opportunityID int, input AddDocumentRequest) (opportunity models.Opportunity, err response.ErrorResponse)
	Query(ctx context.Context, offset, limit int, filters OpportunityQueryRequestParams) (
		opportunities []models.Opportunity, err response.ErrorResponse,
	)
	Count(ctx context.Context, params OpportunityQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateOpportunityRequest, duplicateAllowed bool) (opportunity models.Opportunity, err response.ErrorResponse)
	Update(ctx context.Context, id string, input UpdateOpportunityRequest) (opportunity models.Opportunity, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteOpportunityRequest) (response string, err response.ErrorResponse)

	CheckDuplicate(ctx context.Context, solicitation_number string) (
		bool, models.Opportunity, response.ErrorResponse,
	)

	DownloadAllDocuments(ctx context.Context, id string) (string, response.ErrorResponse)

	ContributionsState(ctx context.Context) (ContributionsStateResponseData, response.ErrorResponse)

	CalculateOpportunityPrice(ctx context.Context, id string) (float64, response.ErrorResponse)

	ImportOpportunities(ctx context.Context, importReq ImportOpportunitiesRequest) response.ErrorResponse
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
