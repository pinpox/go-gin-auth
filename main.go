package main

import (
	"ginjwtauth/controllers"
	"ginjwtauth/middlewares"
	"ginjwtauth/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Load database connection
	models.ConnectDataBase()

	// Load templates
	r.LoadHTMLGlob("views/*")

	// Set up session middleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Set up FlashMessageMiddleware
	r.Use(middlewares.FlashMessageMiddleware())

	// Group routes
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// Register and login routes
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Profile route with JWT authentication middleware
		api.GET("/profile", middlewares.JwtAuthMiddleware(), controllers.Profile)
		api.GET("/user", middlewares.JwtAuthMiddleware(), controllers.Profile)
		api.POST("/user", middlewares.JwtAuthMiddleware(), controllers.UpdateProfile)
	}

	dashboard := r.Group("/dashboard")
	dashboard.Use(middlewares.AuthMiddleware())
	dashboard.GET("/profile", controllers.ProfileIndex)

	web := r.Group("/web")
	{
		web.Use(middlewares.AuthMiddleware())
		web.POST("/user", controllers.UpdateProfile)
	}

	r.GET("/login", controllers.LoginIndex)
	r.GET("/register", controllers.RegisterIndex)

	// New login route that sets a session
	r.POST("/api/loginWithSession", controllers.LoginWithSession)

	// Run the server
	r.Run(":8000")
}
