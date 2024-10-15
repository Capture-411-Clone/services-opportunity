package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M311Department(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Department{})
	if err != nil {
		log.Fatal(err)
	}
}
