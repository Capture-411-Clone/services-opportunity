package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M291Market(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Market{})
	if err != nil {
		log.Fatal(err)
	}
}
