package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M391Opportunity(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Opportunity{})
	if err != nil {
		log.Fatal(err)
	}
}
