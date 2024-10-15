package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M361FiscalYear(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.FiscalYear{})
	if err != nil {
		log.Fatal(err)
	}
}
