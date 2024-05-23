package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RenderTemplate is a custom function to render a template with flash messages
func RenderTemplate(c *gin.Context, templateName string, data gin.H) {
	// Fetch flash messages from the template context
	flashMessages := c.MustGet("getFlashMessages").(func() []string)()

	// Add flash messages to the data to be passed to the template
	data["Flash"] = flashMessages

	// Render the template
	c.HTML(http.StatusOK, templateName, data)
}