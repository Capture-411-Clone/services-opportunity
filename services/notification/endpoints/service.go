package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, input CreateNotificationRequest) ([]models.Notification, response.ErrorResponse)
	Update(ctx context.Context, id int, input UpdateNotificationRequestBody) (models.Notification, response.ErrorResponse)
	Delete(ctx context.Context, input DeleteNotificationRequestBody) (string, response.ErrorResponse)
	Query(
		ctx context.Context, offset, limit int, params NotificationQueryRequestParams,
	) ([]models.Notification, response.ErrorResponse)
	Count(ctx context.Context, params NotificationQueryRequestParams) (int64, response.ErrorResponse)
}

type service struct {
	db       *gorm.DB
	logger   log.Logger
	notifier notificationv2.Notifier
}

func MakeService(db *gorm.DB, logger log.Logger, notifier notificationv2.Notifier) Service {
	return &service{
		db,
		logger,
		notifier,
	}
}
