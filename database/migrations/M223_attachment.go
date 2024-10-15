package migrations

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func M223Attachment(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Attachment{})
	if err != nil {
		log.Fatal(err)
	}
}
