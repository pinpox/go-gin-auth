package controllers

import (
	// "go-gin-auth/models"
	// "go-gin-auth/utils"
	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// "net/http"
	// "strings"
)

func NoteList(c *gin.Context)   {}
func NoteShow(c *gin.Context)   {}
func NoteCreate(c *gin.Context) {}
func NoteUpdate(c *gin.Context) {}
func NoteDelete(c *gin.Context) {}

// ProfileInput represents the input for the profile update endpoint
// type ProfileInput struct {
// 	Name  string `form:"name" binding:"required"`
// 	Email string `form:"email" binding:"required"`
// }
//
// // Profile displays the user's profile
// func ProfileIndex(c *gin.Context) {
// 	session := sessions.Default(c)
// 	username := session.Get("user").(string)
//
// 	user, err := models.GetUserByUsername(username)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
// 		return
// 	}
//
// 	utils.RenderTemplate(c, "profile.tmpl", gin.H{
// 		"Name":  user.Name,
// 		"Email": user.Email,
// 	})
// }
//
// // ProfileUpdate updates the user profile based on the provided input
// func ProfileUpdate(c *gin.Context) {
// 	var (
// 		userID uint
// 		err    error
// 	)
//
// 	session := sessions.Default(c)
// 	username := session.Get("user")
//
// 	// Retrieve user
// 	if username != nil {
// 		user, err := models.GetUserByUsername(username.(string))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
// 			return
// 		}
//
// 		userID = user.ID
// 	}
//
// 	var input ProfileInput
//
// 	// Use ShouldBind with Form instead of ShouldBindQuery
// 	if err := c.ShouldBind(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	u, err := models.GetUserByID(userID)
//
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
// 		return
// 	}
//
// 	u.Name = input.Name
// 	u.Email = input.Email
//
// 	_, saveErr := u.UpdateUser()
//
// 	if saveErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error4": saveErr.Error()})
// 		return
// 	}
//
// 	// Check if the request is from the web (HTML form)
// 	if strings.Contains(c.Request.URL.Path, "/web/") {
// 		// Redirect to a different page
// 		flashMessage := c.MustGet("setFlashMessage").(func(string))
// 		flashMessage("Profile updated successfully")
//
// 		redirectURL := "/dashboard/profile"
// 		c.Redirect(http.StatusSeeOther, redirectURL)
// 	} else {
// 		// Respond with JSON for API requests
// 		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
// 	}
// }
