package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"gorm.io/gorm"
)

type Service interface {
	SendCode(ctx context.Context, input SendCodeRequest) (code string, err response.ErrorResponse)
	Exchange(ctx context.Context, input ExchangeRequest) (sessionCode string, err response.ErrorResponse)
	CheckPhoneExists(ctx context.Context, input CheckPhoneExistsRequest) (exists bool, err response.ErrorResponse)
	CheckEmailExists(ctx context.Context, input CheckEmailExistsRequest) (exists bool, err response.ErrorResponse)
}

type service struct {
	db         *gorm.DB
	logger     log.Logger
	notifierv2 notificationv2.Notifier
}

func MakeService(db *gorm.DB, logger log.Logger, notifierv2 notificationv2.Notifier) Service {
	return &service{
		db,
		logger,
		notifierv2,
	}
}
