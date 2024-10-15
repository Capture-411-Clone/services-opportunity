package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M371SetAside(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.SetAside{})
	if err != nil {
		log.Fatal(err)
	}
}
