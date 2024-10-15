package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M228Verification(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Verification{})
	if err != nil {
		log.Fatal(err)
	}
}
