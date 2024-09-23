package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectToDb() {
	var err error

	dns := os.Getenv("DB_CONNECTION_STRING")
	Connection, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
