package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M229CaptureCost(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.CaptureCost{})
	if err != nil {
		log.Fatal(err)
	}
}
