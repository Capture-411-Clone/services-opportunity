package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M381ContractVehicle(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.ContractVehicle{})
	if err != nil {
		log.Fatal(err)
	}
}
