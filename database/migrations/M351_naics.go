package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M351Naics(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Naics{})
	if err != nil {
		log.Fatal(err)
	}
}
