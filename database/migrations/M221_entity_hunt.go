package migrations

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func M221EntityHunt(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.EntityHunt{})
	if err != nil {
		log.Fatal(err)
	}
}
