package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func FiscalYearsSeeder(db *gorm.DB) {
	fiscalYears := []models.FiscalYear{

		{
			UserID: 1,
			Year:   "2024",
		},

		{
			UserID: 1,
			Year:   "2023",
		},

		{
			UserID: 1,
			Year:   "2022",
		},

		{
			UserID: 1,
			Year:   "2021",
		},

		{
			UserID: 1,
			Year:   "2020",
		},

		{
			UserID: 1,
			Year:   "2019",
		},

		{
			UserID: 1,
			Year:   "2018",
		},

		{
			UserID: 1,
			Year:   "2017",
		},

		{
			UserID: 1,
			Year:   "2016",
		},

		{
			UserID: 1,
			Year:   "2015",
		},

		{
			UserID: 1,
			Year:   "2014",
		},

		{
			UserID: 1,
			Year:   "2013",
		},

		{
			UserID: 1,
			Year:   "2012",
		},

		{
			UserID: 1,
			Year:   "2011",
		},

		{
			UserID: 1,
			Year:   "2010",
		},
	}

	err := db.Create(&fiscalYears).Error
	if err != nil {
		log.Fatal(err)
	}
}
