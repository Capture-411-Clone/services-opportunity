package verify

import (
	"path/filepath"

	"github.com/Capture-411/services-opportunity/kit/log"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"github.com/Capture-411/services-opportunity/services/verify/endpoints"
	"github.com/Capture-411/services-opportunity/services/verify/transports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(r *echo.Echo, db *gorm.DB, notifierv2 notificationv2.Notifier, logger log.Logger, prefix string) {
	service := endpoints.MakeService(db, logger, notifierv2)
	transports.RegisterHandlers(r, service, logger, filepath.Join("/api", prefix))
}
