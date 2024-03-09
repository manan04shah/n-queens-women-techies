package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connectDb() {
	db, err := gorm.Open(postgres.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the database")
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	//TODO: Add Migrations

	Database = DbInstance{Db: db}
}
