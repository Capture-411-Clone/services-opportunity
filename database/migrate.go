package database

import (
	"fmt"
	"log"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/database/migrations"
	"gorm.io/gorm"
)

func Confirm(action string) bool {
	var response string
	log.Println("Are you sure you want to do " + action + " in the database? (yes/no)")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response == "yes"
}

func Migrate(db *gorm.DB) {
	if config.AppConfig.Environment != "development" {
		// prompt user to confirm
		if !Confirm("migrate") {
			log.Fatal("Migration aborted")
		}
	}

	// migrate the dependent many2many before the main table
	migrations.M181SiteInfo(db)
	migrations.M201RegisterInvite(db)
	migrations.M211Invite(db)
	migrations.M223Attachment(db)
	migrations.M231Role(db)
	migrations.M221EntityHunt(db)
	migrations.M222User(db)
	migrations.M228Verification(db)
	migrations.M229CaptureCost(db)
	migrations.M230PassiveRevenue(db)
	migrations.M291Market(db)
	migrations.M311Department(db)
	migrations.M321Agency(db)
	migrations.M331Office(db)
	migrations.M351Naics(db)
	migrations.M361FiscalYear(db)
	migrations.M371SetAside(db)
	migrations.M381ContractVehicle(db)
	migrations.M372Document(db)
	migrations.M391Opportunity(db)
	migrations.M411Wishlist(db)
	migrations.M511Order(db)
	migrations.M611Notification(db)
}
