package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M611Notification(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatal(err)
	}
}
