package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDB() {
	domainNameServer := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(domainNameServer), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ❌", err)
	}

	Database = database
	log.Println("Database Connected Successfully! ✅")

}
