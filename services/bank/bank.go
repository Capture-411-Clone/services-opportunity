package bank

import (
	"path/filepath"

	"github.com/Capture-411/services-opportunity/kit/log"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/Capture-411/services-opportunity/services/bank/transports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(r *echo.Echo, db *gorm.DB, logger log.Logger, notifier notificationv2.Notifier, prefix string) {
	service := endpoints.MakeService(db, logger, notifier)
	transports.RegisterHandlers(r, service, logger, filepath.Join("/api", prefix))
}
