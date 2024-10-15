package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Check(ctx context.Context, code string) (invite models.Invite, err response.ErrorResponse)
	Query(ctx context.Context, offset, limit int, filters InviteQueryRequestParams) (invites []models.Invite, err response.ErrorResponse)
	Count(ctx context.Context, params InviteQueryRequestParams) (count int64, err response.ErrorResponse)
	Create(ctx context.Context, input CreateInviteRequest) (invite models.Invite, err response.ErrorResponse)
	Update(ctx context.Context, id string, input UpdateInviteRequest) (invite models.Invite, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteInviteRequest) (response []int, err response.ErrorResponse)
	// No transport
	Use(ctx context.Context, code string) (invite models.Invite, err response.ErrorResponse)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{db, logger}
}
