package main

import (
	"log"
	"os"

	"go-gin-auth/controllers"
	"go-gin-auth/middlewares"
	"go-gin-auth/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/joho/godotenv"
)

func loadEnv() {

	required := []string{
		"SESSION_SECRET",
		"ADMIN_USERNAME",
		"ADMIN_PASSWORD",
		"ADMIN_NAME",
		"ADMIN_EMAIL",
		"LISTEN_ADDRESS",
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file provided, checking environment")
	}

	for _, v := range required {
		if len(os.Getenv(v)) == 0 {
			log.Fatalf("FATAL: Environment var: %s is not set\n", v)
		}
	}
}

func main() {

	loadEnv()

	// Initialize the Gin router
	r := gin.Default()

	// Load database connection
	models.ConnectDataBase()

	// Create superuser account from env
	var superuser = models.User{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Name:     os.Getenv("ADMIN_NAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
	}

	_, err := superuser.SaveAdmin()
	if err != nil {
		panic(err)
	}

	// Load templates
	// TODO use embed
	r.LoadHTMLGlob("views/*")

	// Global middlewares

	// Sessions
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	r.Use(sessions.Sessions("session", store))

	// Log to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovers panics, writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Login, Logout and register
	r.GET("/", controllers.Index)

	r.GET("/login", controllers.LoginIndex)
	r.GET("/register", controllers.RegisterIndex)
	r.GET("/logout", controllers.Logout)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Authorized group
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	authorized.GET("/user", controllers.ProfileIndex)
	authorized.POST("/user", controllers.ProfileUpdate)

	{
		note := authorized.Group("/note")
		note.GET("/", controllers.NoteList)

		note.GET("/create", controllers.NoteCreateShow)
		note.POST("/create", controllers.NoteCreate)

		note.GET("/:id", controllers.NoteShow)
		note.PUT("/:id", controllers.NoteUpdate)
		note.DELETE("/:id", controllers.NoteDelete)
	}

	// Run the server
	r.Run(os.Getenv("LISTEN_ADDRESS"))
}
