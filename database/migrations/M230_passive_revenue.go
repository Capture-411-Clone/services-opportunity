package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M230PassiveRevenue(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.PassiveRevenue{})
	if err != nil {
		log.Fatal(err)
	}
}
