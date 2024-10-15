package database

import (
	"fmt"
	"log"
	"os"

	"go.com/railwayticket/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetUpDatabase() {

	//Build the Connection string using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	//connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable Info level logging
	})

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	//Migrate the passenger model to the database
	DB.AutoMigrate(&models.Passenger{})
}
