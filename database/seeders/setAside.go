package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func SetAsidesSeeder(db *gorm.DB) {
	setAsides := []models.SetAside{

		{
			UserID: 1,
			Name:   "8(a) Sole Source",
		},

		{
			UserID: 1,
			Name:   "Economically Disadvantaged Women-Owned Small Business",
		},

		{
			UserID: 1,
			Name:   "HUBZone",
		},

		{
			UserID: 1,
			Name:   "Service-Disabled Veteran-Owned Small Business",
		},

		{
			UserID: 1,
			Name:   "Small Business Set-Aside - Partial",
		},

		{
			UserID: 1,
			Name:   "Small Business Set-Aside - Total",
		},

		{
			UserID: 1,
			Name:   "Unrestricted",
		},

		{
			UserID: 1,
			Name:   "Women-Owned Small Business",
		},
	}

	err := db.Create(&setAsides).Error
	if err != nil {
		log.Fatal(err)
	}
}
