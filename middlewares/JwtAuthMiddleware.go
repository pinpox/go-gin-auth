// middlewares/JwtAuthMiddleware.go
package middlewares

import (
	"strings"
	"os"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"ginjwtauth/utils"
	"net/http"
)

var secretKey = os.Getenv("SECRET_KEY")

// JwtAuthMiddleware is a middleware for JWT authentication
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Verify the token
		token, err := utils.ExtractTokenID(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID in the context for further use in the handler
		c.Set("userID", token)

		c.Next()
	}
}

// TokenCookie sets the JWT token in a cookie
func TokenCookie(c *gin.Context) {
	// Get the token from the request header
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Set the token in the cookie
	c.SetCookie("token", tokenString, 0, "/", "localhost", false, true)
}

// ExtractTokenID extracts the user ID from the JWT token
func ExtractTokenID(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, fmt.Errorf("Invalid token")
}
