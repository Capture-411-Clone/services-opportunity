package migrations

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func M211Invite(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Invite{})

	if err != nil {
		log.Fatal(err)
	}
}
