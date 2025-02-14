package middleware

import (
	"app/internal/auth"
	contextKeys "app/internal/contextkeys"
	"app/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// assign the value of the cookie which is the token to the token variable
		tokenString := cookie

		// Validate the token
		claims, err := auth.ValidateJWTAndExtractClaims(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract user ID from claims (assuming "sub" holds the user ID as a string)
		sub, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, err := strconv.ParseInt(sub, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Get the user from the database using the user ID
		user, err := services.GetUserByID(db, userID)
		if err != nil || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Store the user in context using our defined key.
		c.Set(string(contextKeys.UserKey), user)

		// Token is valid, proceed to the next handler
		c.Next()
	}
}
