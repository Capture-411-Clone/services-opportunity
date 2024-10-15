package contractVehicle

import (
	"path/filepath"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/services/contractVehicle/endpoints"
	"github.com/Capture-411/services-opportunity/services/contractVehicle/transports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(r *echo.Echo, db *gorm.DB, logger log.Logger, prefix string) {
	service := endpoints.MakeService(db, logger)
	transports.RegisterHandlers(r, service, logger, filepath.Join("/api", prefix))
}
