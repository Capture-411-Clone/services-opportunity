package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func MarketSeeder(db *gorm.DB) {
	marketSeeders := []models.Market{
		{
			// ID: 1
			Name:   "Defense",
			UserID: 1,
		},
		{
			// ID : 2
			Name:   "Civilian",
			UserID: 1,
		},
	}

	err := db.Create(&marketSeeders).Error
	if err != nil {
		log.Fatal(err)
	}
}
