// controllers/AuthController.go
package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"ginjwtauth/models"
	"ginjwtauth/utils"
	"net/http"
	"github.com/gin-contrib/sessions"
)

// RegisterIndex displays the registration form
func RegisterIndex(c *gin.Context) {
	utils.RenderTemplate(c, "register.tmpl", gin.H{
		"title": "Register",
	})
}

type RegisterInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
}

// Register handles user registration
func Register(c *gin.Context) {
	var input RegisterInput

	// Bind and validate input
	if err := c.ShouldBind(&input); err != nil {

		log.Println("NOT Parsing")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})


		return
	}

	// Create a new user
	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Email:    input.Email,
	}

	// Save the user to the database
	_, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login
func Login(c *gin.Context) {
	var input LoginInput

	log.Println("Im in login")

	// Bind and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user credentials and generate a token
	token, err := models.LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Set the token in a cookie
	utils.TokenCookie(c)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Profile handles user profile
func Profile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Retrieve user from the database
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user details
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
	})
}

// Login Index Page
func LoginIndex (c *gin.Context) {
	utils.RenderTemplate(c, "login.tmpl", gin.H{
		"title": "Login",
	})
}

// LoginWithSession handles user login and sets a session
func LoginWithSession(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Authenticate user
	token, err := models.LoginCheck(username, password)

	if err != nil {
		// Authentication failed, return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Set a session variable
	session := sessions.Default(c)
	session.Set("token", token)
	session.Set("user", username)
	session.Save()

	redirectURL := "/dashboard/profile"
    c.Redirect(http.StatusSeeOther, redirectURL)
}
