package migrations

import (
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func M411Wishlist(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&models.Wishlist{})
	if err != nil {
		log.Fatal(err)
	}
}
