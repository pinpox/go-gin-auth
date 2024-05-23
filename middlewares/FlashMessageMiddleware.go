package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// FlashMessageMiddleware adds flash message functionality to the Gin context
func FlashMessageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Function to set flash messages
		c.Set("setFlashMessage", func(message string) {
			session := sessions.Default(c)
			session.AddFlash(message)
			session.Save()
		})

		// Function to get flash messages
		c.Set("getFlashMessages", func() []string {
			session := sessions.Default(c)
			flashes := session.Flashes()
			messages := make([]string, len(flashes))
			for i, flash := range flashes {
				messages[i] = flash.(string)
			}
			session.Save()
			return messages
		})

		c.Next()
	}
}