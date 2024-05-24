package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/models"
	"go-gin-auth/utils"
	"log"
	"net/http"
)

func NoteList(c *gin.Context) {

	user, err := GetCurrentUser(c)
	if err != nil {
		utils.ErrorRedirect(c, err, "No user found", "/")
		return
	}

	notes, err := models.GetNotesByUser(user.ID)

	if err != nil {
		utils.ErrorRedirect(c, err, "Failed to get notes", "/")
		return
	}

	utils.RenderTemplate(c, "note-list.tmpl", gin.H{"notes": notes})

}
func NoteShow(c *gin.Context) {}

// PUT
func NoteUpdate(c *gin.Context) {

	var input NoteInput

	// Bind and validate input
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}

	session := sessions.Default(c)
	username := session.Get("user").(string)

	user, err := models.GetUserByUsername(username)
	if err != nil {
		utils.ErrorRedirect(c, err, "User not found", "/")
		return
	}

	n, err := models.GetNoteByID(1)
	if err != nil || n.UserID != int(user.ID) {
		utils.ErrorRedirect(c, err, "Failed to update note", "/")
		return
	}

	n.Text = input.Text
	n.Title = input.Title

	if err = n.UpdateNote(); err != nil {
		utils.ErrorRedirect(c, err, "Failed to update note", "/")
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

	utils.FlashSuccess(c, "Note deleted successfully")
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
		log.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	session := sessions.Default(c)
	username := session.Get("user").(string)

	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
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
		utils.ErrorRedirect(c, err, "Failed to save note", "/")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Node created successfully"})

}
