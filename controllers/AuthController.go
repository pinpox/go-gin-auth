package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/models"
	"go-gin-auth/utils"
	"log"
	"net/http"
)

// RegisterIndex displays the registration form
func RegisterIndex(c *gin.Context) {
	utils.RenderTemplate(c, "register-form", gin.H{})
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
		log.Println(err)
		utils.FlashSuccess(c, "Invalid registration data provided")
		c.Redirect(http.StatusBadRequest, "/")
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
		log.Println(err)
		utils.FlashSuccess(c, "Failed to save user")
		c.Redirect(http.StatusBadRequest, "/")
		return
	}

	utils.FlashSuccess(c, "User registered successfully")
	c.Redirect(http.StatusSeeOther, "/")
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginIndex(c *gin.Context) { utils.RenderTemplate(c, "login-form", gin.H{"title": "Login"}) }

// Logout handles user login and sets a session
func Logout(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("No user to log out, invalid session token")
		return
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		log.Println(err)
		utils.FlashSuccess(c, "Logout failed")
		c.Redirect(http.StatusInternalServerError, "/")
		return
	}

	utils.FlashInfo(c, "Logged out successfully")
	c.Redirect(http.StatusSeeOther, "/")
}

// Login handles user login and sets a session
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Authenticate user
	err := models.LoginCheck(username, password)

	if err != nil {
		utils.FlashError(c, "Invalid username or password")
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// Set a session variable
	session := sessions.Default(c)
	session.Set("user", username)
	session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}
