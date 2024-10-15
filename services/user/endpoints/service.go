package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(
		ctx context.Context, offset, limit int, params UserQueryRequestParams,
	) ([]models.User, response.ErrorResponse)
	Count(ctx context.Context, filters UserQueryRequestParams) (int64, response.ErrorResponse)
	Create(ctx context.Context, req CreateUserRequest) (models.User, response.ErrorResponse)
	Update(ctx context.Context, id string, req UpdateUserRequest) (models.User, response.ErrorResponse)
	Delete(ctx context.Context, input DeleteUserRequest) (response string, err response.ErrorResponse)
	UpdateAccount(ctx context.Context, req UpdateUserAccountRequest) (models.User, response.ErrorResponse)
	Suspend(ctx context.Context, id string) (models.User, response.ErrorResponse)
	ToggleVerifyEmail(ctx context.Context, id string) (models.User, response.ErrorResponse)
	ToggleVerifyPhone(ctx context.Context, id string) (models.User, response.ErrorResponse)
	ApprovePolicy(ctx context.Context) (models.User, response.ErrorResponse)
	CheckIsUniqueField(ctx context.Context, fields, value string) (exists bool, user models.User, err response.ErrorResponse)
	StaffReport(ctx context.Context, offset, limit int, params StaffReportQueryRequestParams) ([]StaffReportResponseData, response.ErrorResponse)
	StaffReportCount(ctx context.Context, params StaffReportQueryRequestParams) (int64, response.ErrorResponse)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{
		db,
		logger,
	}
}
