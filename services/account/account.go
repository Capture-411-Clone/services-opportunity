package account

import (
	"path/filepath"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/labstack/echo/v4"

	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"github.com/Capture-411/services-opportunity/services/account/endpoints"
	http "github.com/Capture-411/services-opportunity/services/account/transports"
	"gorm.io/gorm"
)

func Register(r *echo.Echo, db *gorm.DB, logger log.Logger, notifierv2 notificationv2.Notifier, AccessTokenSigningKey string, RefreshTokenSigningKey string, AccessTokenTokenExpiration int, RefreshTokenExpiration int, prefix string) {
	service := endpoints.MakeService(db, logger, notifierv2, AccessTokenSigningKey, RefreshTokenSigningKey, AccessTokenTokenExpiration, RefreshTokenExpiration)
	http.RegisterHandlers(r, service, logger, filepath.Join("/api", prefix))
}
