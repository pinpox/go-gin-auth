package controllers

import (
	"go-gin-auth/models"
	"go-gin-auth/utils"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strings"
)

func NoteList(c *gin.Context)   {}
func NoteShow(c *gin.Context)   {}
func NoteUpdate(c *gin.Context) {}
func NoteDelete(c *gin.Context) {}

func NoteCreateShow(c *gin.Context) {
	log.Println("Rendireng note page")
	utils.RenderTemplate(c, "newnote.tmpl", gin.H{})
}

// NoteInput represents the input for the note create/update endpoints
type NoteInput struct {
	Title string `form:"title" binding:"required"`
	Text  string `form:"text" binding:"required"`
}

func NoteCreate(c *gin.Context) {

	var input NoteInput

	// Bind and validate input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	session := sessions.Default(c)
	username := session.Get("user").(string)

	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}


	// Create a new note
	note := models.Note{
		Title: input.Title,
		Text:  input.Text,
		User: user,
	}

	// Save the user to the database
	_, err = note.SaveNote()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

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
