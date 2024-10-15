package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M511Order(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal(err)
	}
}
