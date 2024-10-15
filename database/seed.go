package database

import (
	"log"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/database/seeders"
	"gorm.io/gorm"
)

func SeedAllTables(db *gorm.DB) {
	if config.AppConfig.Environment != "development" {
		// prompt user to confirm
		if !Confirm("seed") {
			log.Fatal("Seed aborted")
		}
	}

	seeders.SiteInfoSeeder(db)
	seeders.RoleSeeder(db)
	seeders.UserSeeder(db)
	seeders.MarketSeeder(db)
	seeders.DepartmentsSeeder(db)
	seeders.AgenciesSeeder(db)
	seeders.OfficeSeeder(db)
	seeders.NaicsesSeeder(db)
	seeders.FiscalYearsSeeder(db)
	seeders.SetAsidesSeeder(db)
	seeders.ContractVehiclesSeeder(db)
	seeders.OpportunitySeeder(db)
}
