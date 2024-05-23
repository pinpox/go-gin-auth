package models

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// ConnectDataBase connects to the database using environment variables
func ConnectDataBase() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database

	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate creates or updates database tables based on model definitions
	DB.AutoMigrate(&User{})
}
