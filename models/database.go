package models

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDataBase connects to the database using environment variables
func ConnectDataBase() {
	// Load environment variables from .env file

	var err error


	// Connect to the database

	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate creates or updates database tables based on model definitions
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Note{})
}
