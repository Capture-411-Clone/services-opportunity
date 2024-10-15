package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func SiteInfoSeeder(db *gorm.DB) {
	siteInfo := []models.SiteInfo{
		{
			CompanyName:      "Capture411",
			MissionStatement: "Your Mission Statement",
			History:          "Your History",
			Goal:             "Your Goals",
			Value:            "Your Values",
			Achievement:      "Your Achievements",
			Member:           "Your Member",
			GeneralContact:   "General Contact Info",
			Address:          "P.O. Box 1074, Dumfries, VA 22026 ",
			SocialMedia:      "instagram.com/your_business, telegram.com/you_business",
			PhoneNumber:      "123-456-7890",
			EmailAddress:     "info@capture411.com",
			OfficeHours:      "Monday-Friday 9 AM - 5 PM",
			PhysicalAddress:  "Your Company Physical Address",
			MapOrDirections:  "Map or Directions Link",
			EmergencyContact: "Emergency Contact Info",
			FeedbackLink:     "Feedback Link",
			SupportLink:      "Support Link",
		},
	}

	err := db.Create(&siteInfo).Error
	if err != nil {
		log.Fatal(err)
	}
}
