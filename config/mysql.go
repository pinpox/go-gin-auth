// // config/mysql.go
// package config
//
// import (
// 	"fmt"
// 	"log"
// 	"os"
//
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/joho/godotenv"
// 	"ginjwtauth/models"
// )
//
// var DB *gorm.DB
//
// // ConnectDataBase connects to the database using environment variables
// func ConnectDataBase() {
// 	// Load environment variables from .env file
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
//
// 	// Retrieve database connection details from environment variables
// 	Dbdriver := os.Getenv("DB_DRIVER")
// 	DbHost := os.Getenv("DB_HOST")
// 	DbUser := os.Getenv("DB_USER")
// 	DbPassword := os.Getenv("DB_PASSWORD")
// 	DbName := os.Getenv("DB_NAME")
// 	DbPort := os.Getenv("DB_PORT")
//
// 	// Construct the database URL
// 	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
//
// 	// Connect to the database
// 	DB, err = gorm.Open(Dbdriver, DBURL)
// 	if err != nil {
// 		fmt.Println("Cannot connect to database ", Dbdriver)
// 		log.Fatal("connection error:", err)
// 	} else {
// 		fmt.Println("We are connected to the database ", Dbdriver)
// 	}
//
// 	// AutoMigrate creates or updates database tables based on model definitions
// 	DB.AutoMigrate(&models.User{})
// }
