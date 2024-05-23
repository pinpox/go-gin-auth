package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

// AuthMiddleware is a middleware to check session for authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		log.Println(c.Get("userID"))

		// Check if the "user" key is present in the session
		user := session.Get("user")
		if user == nil {
			log.Println("No user found in AuthMiddleware, redirecting to login")
			// Redirect to the login page if not authenticated
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // Stop the execution of subsequent middleware and the handler
			return
		}

		// Continue with the request if authenticated
		c.Next()
	}
}
