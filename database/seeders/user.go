package seeders

import (
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/faker"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func UserSeeder(db *gorm.DB) {
	users := []models.User{
		{
			Name:         "Christina",
			ReferralCode: "BOSSLADY",
			Phone: dtp.NullString{
				String: "7203511106",
				Valid:  true,
			},
			LastName:        "Mee",
			Password:        "$2a$14$ALuVvY5tXovI9WFyxed/eORguxkDhWaEgHARuOMdekOloWFL2zuMO", // password
			PhoneVerifiedAt: faker.SQLNow(),
			Username: dtp.NullString{
				String: "bosslady",
				Valid:  true,
			},
			Email: dtp.NullString{
				String: "christina.mee@gmail.com",
				Valid:  true,
			},
			EmailVerifiedAt: faker.SQLNow(),
			Roles: []*models.Role{
				{ID: 1},
			},
		},
		{
			Name:         "Matthew",
			ReferralCode: "mattgoz",
			Phone: dtp.NullString{
				String: "720-931-1196",
				Valid:  true,
			},
			LastName:        "L.Gonzales",
			Password:        "$2a$14$ALuVvY5tXovI9WFyxed/eORguxkDhWaEgHARuOMdekOloWFL2zuMO", // password
			PhoneVerifiedAt: faker.SQLNow(),
			IDCode: dtp.NullString{
				String: "008-84-0497",
				Valid:  true,
			},
			Username: dtp.NullString{
				String: "mattgoz",
				Valid:  true,
			},
			Email: dtp.NullString{
				String: "matt.gonzales@gmail.com",
				Valid:  true,
			},
			EmailVerifiedAt:    faker.SQLNow(),
			ProfileCompletedAt: faker.SQLNow(),
			SuspendedAt:        dtp.NullTime{},
			Roles: []*models.Role{
				{ID: 3},
			},
		},
	}

	err := db.Create(&users).Error
	if err != nil {
		log.Fatal(err)
	}
}
