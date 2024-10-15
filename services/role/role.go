package role

import (
	"path/filepath"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(r *echo.Echo, db *gorm.DB, logger log.Logger, prefix string) {
	service := MakeService(db, logger)
	RegisterHandlers(r, service, logger, filepath.Join("/api", prefix))
}
