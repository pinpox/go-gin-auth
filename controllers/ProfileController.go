package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/models"
	"go-gin-auth/utils"
	"log"
	"net/http"
)

// ProfileInput represents the input for the profile update endpoint
type ProfileInput struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

// Profile displays the user's profile
func ProfileIndex(c *gin.Context) {

	user, err := GetCurrentUser(c)

	if err != nil {
		log.Println(err)
		utils.FlashError(c, "User not found!")
		c.Redirect(http.StatusSeeOther, "/user")
		return
	}

	utils.RenderTemplate(c, "profile.tmpl", gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}

// ProfileUpdate updates the user profile based on the provided input
func ProfileUpdate(c *gin.Context) {


	currentUser, err := GetCurrentUser(c)
	if err != nil {
		log.Println(err)
		utils.FlashError(c, "Failed to update user!")
		c.Redirect(http.StatusSeeOther, "/user")
		c.Abort()
		return
	}

	var input ProfileInput

	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		utils.FlashError(c, "Failed to update user!")
		c.Redirect(http.StatusSeeOther, "/user")
		c.Abort()
		return
	}

	u, err := models.GetUserByID(currentUser.ID)

	if err != nil {
		log.Println(err)
		utils.FlashError(c, "Failed to update user!")
		c.Redirect(http.StatusSeeOther, "/user")
		c.Abort()
		return
	}

	u.Name = input.Name
	u.Email = input.Email

	_, saveErr := u.UpdateUser()

	if saveErr != nil {
		log.Println(saveErr)
		utils.FlashError(c, "Failed to update user!")
		c.Redirect(http.StatusSeeOther, "/user")
		c.Abort()
		return
	} else {
		utils.FlashError(c, "Profile updated successfully")
		utils.RenderTemplate(c, "profile-form", gin.H{"Name": u.Name, "Email": u.Email})
	}
}

// GetCurrentUser retrieves the current user of a context
func GetCurrentUser(c *gin.Context) (models.User, error) {

	var user models.User
	var err error

	session := sessions.Default(c)
	username := session.Get("user")

	if username != nil {
		return models.GetUserByUsername(username.(string))
	} else {
		return user, err
	}
}
