package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M331Office(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Office{})
	if err != nil {
		log.Fatal(err)
	}
}
