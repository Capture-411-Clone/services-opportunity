package database

import (
	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"

	"log"
)

func DropJoinTables(db *gorm.DB) {
	// TODO: handle errors
	err := db.Migrator().DropTable("user_permission")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable("user_role")
	if err != nil {
		log.Fatal(err)
	}
}

func DropTables(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.EntityHunt{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.SiteInfo{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.RegisterInvite{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Invite{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Attachment{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.CaptureCost{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.PassiveRevenue{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Verification{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Role{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.Market{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.Department{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.Agency{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.Office{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.Naics{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.FiscalYear{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.SetAside{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.ContractVehicle{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Document{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Opportunity{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Wishlist{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Order{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Notification{})
	if err != nil {
		log.Fatal(err)
	}
}

func DropAll(db *gorm.DB) {
	if config.AppConfig.Environment != "development" {
		if !Confirm("drop all tables") {
			log.Fatal("Migration aborted")
		}
	}

	DropJoinTables(db.Debug())
	DropTables(db.Debug())
}
