package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M181SiteInfo(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.SiteInfo{})
	if err != nil {
		log.Fatal(err)
	}
}
