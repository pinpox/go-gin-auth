package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/models"
	"go-gin-auth/utils"
	"log"
	"net/http"
	// "strings"
)

func NoteList(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("user").(string)

	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	notes, err := models.GetNotesByUser(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get notes"})
		return
	}

	// c.JSON(http.StatusFound, gin.H{"notes": notes})
	utils.RenderTemplate(c, "note-list.tmpl", gin.H{"notes": notes})

}
func NoteShow(c *gin.Context) {}

// PUT
func NoteUpdate(c *gin.Context) {

	log.Println("updating note now")

	var input NoteInput

	// Bind and validate input
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err.Error())
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



	n, err := models.GetNoteByID(1)
	if err != nil || n.UserID != int(user.ID) {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Editing note not possible"})
		c.Abort()
		return
	}

	n.Text = input.Text
	n.Title = input.Title


	if err = n.UpdateNote(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.RenderTemplate(c, "note-partial.tmpl", gin.H{
		"ID":    n.ID,
		"Title": n.Title,
		"Text":  n.Text,
	})
}

func NoteDelete(c *gin.Context) {
	// Example on how to set flashMessage to redirect after deletion
	// TODO actually delete
	flashMessage := c.MustGet("setFlashMessage").(func(string))
	flashMessage("Note Deleted successfully")
	c.Redirect(http.StatusSeeOther, "/note")

}

func NoteCreateShow(c *gin.Context) {
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
		User:  user,
	}

	// Save the user to the database
	_, err = note.SaveNote()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Node created successfully"})

}

//
// // Profile displays the user's profile
// func ProfileIndex(c *gin.Context) {
//	session := sessions.Default(c)
//	username := session.Get("user").(string)
//
//	user, err := models.GetUserByUsername(username)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
//		return
//	}
//
//	utils.RenderTemplate(c, "profile.tmpl", gin.H{
//		"Name":  user.Name,
//		"Email": user.Email,
//	})
// }
//
// // ProfileUpdate updates the user profile based on the provided input
// func ProfileUpdate(c *gin.Context) {
//	var (
//		userID uint
//		err    error
//	)
//
//	session := sessions.Default(c)
//	username := session.Get("user")
//
//	// Retrieve user
//	if username != nil {
//		user, err := models.GetUserByUsername(username.(string))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
//			return
//		}
//
//		userID = user.ID
//	}
//
//	var input ProfileInput
//
//	// Use ShouldBind with Form instead of ShouldBindQuery
//	if err := c.ShouldBind(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	u, err := models.GetUserByID(userID)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
//		return
//	}
//
//	u.Name = input.Name
//	u.Email = input.Email
//
//	_, saveErr := u.UpdateUser()
//
//	if saveErr != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error4": saveErr.Error()})
//		return
//	}
//
//	// Check if the request is from the web (HTML form)
//	if strings.Contains(c.Request.URL.Path, "/web/") {
//		// Redirect to a different page
//		flashMessage := c.MustGet("setFlashMessage").(func(string))
//		flashMessage("Profile updated successfully")
//
//		redirectURL := "/dashboard/profile"
//		c.Redirect(http.StatusSeeOther, redirectURL)
//	} else {
//		// Respond with JSON for API requests
//		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
//	}
// }
