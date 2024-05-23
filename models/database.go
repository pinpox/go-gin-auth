package models

import (
	"log"


	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
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
	DB.AutoMigrate(&Note{})
}

func SeedData(input any) {
	DB.Save(&input)
}
