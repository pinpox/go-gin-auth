package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/models"
	"go-gin-auth/utils"
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
		utils.ErrorRedirect(c, err, "User not found", "/")
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
		utils.ErrorRedirect(c, err, "Failed to update user", "/user")
		return
	}

	var input ProfileInput

	if err := c.ShouldBind(&input); err != nil {
		utils.ErrorRedirect(c, err, "Failed to update user", "/user")
		return
	}

	u, err := models.GetUserByID(currentUser.ID)

	if err != nil {
		utils.ErrorRedirect(c, err, "Failed to update user", "/user")
		return
	}

	u.Name = input.Name
	u.Email = input.Email

	_, saveErr := u.UpdateUser()

	if saveErr != nil {
		utils.ErrorRedirect(c, err, "Failed to update user", "/user")
		return

	} else {
		utils.FlashSuccess(c, "Profile updated successfully")
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
