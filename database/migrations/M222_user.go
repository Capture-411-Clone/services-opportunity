package migrations

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func M222User(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
