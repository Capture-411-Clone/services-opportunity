package database

import (
	"fmt"
	"log"

	"github.com/Capture-411/services-opportunity/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	var (
		host     = config.AppConfig.DatabaseHost
		username = config.AppConfig.DatabaseUsername
		password = config.AppConfig.DatabasePassword
		dbName   = config.AppConfig.DatabaseName
		port     = config.AppConfig.DatabasePort
		sslMode  = config.AppConfig.DatabaseSslMode
	)

	cert := ""
	if config.AppConfig.DatabaseSslMode != "disable" {
		cert = "sslrootcert=" + config.AppConfig.DatabaseSslRootCert
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d"+
		" sslmode=%s "+cert+" TimeZone=America/New_York",
		host, username, password, dbName, port, sslMode)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Execute the SQL command to create the pg_trgm extension if it doesn't exist
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;").Error
	if err != nil {
		log.Fatalf("failed to create pg_trgm extension: %v", err)
	}

	err = db.Exec("SELECT set_limit(0.2);").Error // Adjust 0.1 to your desired threshold
	if err != nil {
		log.Fatalf("failed to set trigram similarity threshold: %v", err)
	}

	return db
}
