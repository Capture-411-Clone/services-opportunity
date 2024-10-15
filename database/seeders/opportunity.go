package seeders

import (
	"log"
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func OpportunitySeeder(db *gorm.DB) {
	opportunitySeeders := []models.Opportunity{
		{
			// ID: 1
			Title:       "Opportunity Sample",
			Description: "Are you ready to be a part of an exciting and dynamic team? We are currently offering a unique opportunity that could shape your career and make a significant impact on our organization. Join us in our quest for excellence as we work on innovative projects and contribute to our success.",
			UserID:      1,
			ApprovedAt: dtp.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			Market:                 "Market Sample",
			Department:             "Department Sample",
			Agency:                 "Agency Sample",
			Office:                 "Office Sample",
			SolicitationNumber:     "50000P",
			UserContractValue:      "800000",
			UserKnowsContractValue: true,
			CrawlerContractValue:   "",
			Naics:                  "5411",
			FiscalYear:             "2020",
			SetAside:               "Set Aside Sample",
			ContractVehicle:        "Contract Vehicle Sample",
		},
		{
			// ID: 2
			Title:       "Innovation challenge",
			Description: "Join our innovation challenge! This is a great chance to showcase your skills and make a real impact on our company. We're looking for creative thinkers who are ready to take on a challenge.",
			UserID:      2,
			ApprovedAt: dtp.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			Market:                 "Market Sample",
			Department:             "Department Sample",
			Agency:                 "Agency Sample",
			Office:                 "Office Sample",
			SolicitationNumber:     "50000P",
			UserContractValue:      "800000",
			UserKnowsContractValue: true,
			CrawlerContractValue:   "",
			Naics:                  "5411",
			FiscalYear:             "2020",
			SetAside:               "Set Aside Sample",
			ContractVehicle:        "Contract Vehicle Sample",
		},
	}

	err := db.Create(&opportunitySeeders).Error
	if err != nil {
		log.Fatal(err)
	}
}
