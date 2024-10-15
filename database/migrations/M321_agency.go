package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M321Agency(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Agency{})
	if err != nil {
		log.Fatal(err)
	}
}
