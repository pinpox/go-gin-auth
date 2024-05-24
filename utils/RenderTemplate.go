package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FlashInfo(c *gin.Context, message string) {
	flash(c, message, "info")
}

func FlashSuccess(c *gin.Context, message string) {
	flash(c, message, "success")
}

func FlashError(c *gin.Context, message string) {
	flash(c, message, "error")
}

func flash(c *gin.Context, message, flashtype string) {
	session := sessions.Default(c)
	session.AddFlash(message, flashtype)
	session.Save()

}

// RenderTemplate is a custom function to render a template with flash messages
func RenderTemplate(c *gin.Context, templateName string, data gin.H) {

	// Get flash mesages from context and add them to the data to be passed to
	// the template

	session := sessions.Default(c)

	data["FlashError"] = getFlashes(c, "error")
	data["FlashInfo"] = getFlashes(c, "info")
	data["FlashSuccess"] = getFlashes(c, "success")

	session.Save()

	// Render the template
	c.HTML(http.StatusOK, templateName, data)
}

func getFlashes(c *gin.Context, flashtype string) []string {

	session := sessions.Default(c)
	flashes := session.Flashes(flashtype)
	messages := make([]string, len(flashes))
	for i, flash := range flashes {
		messages[i] = flash.(string)
	}

	session.Save()
	return messages
}
