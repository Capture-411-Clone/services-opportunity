package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M372Document(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Document{})
	if err != nil {
		log.Fatal(err)
	}
}
