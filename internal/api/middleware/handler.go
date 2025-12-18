package api

import (
	"net/http"

	"todoDB/internal/infra/auth_jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware parses and validates JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		authClaims, err := auth_jwt.Par(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// attach user info to context
		c.Set("userID", userID)
		c.Next()
	}
}

// FromContext retrieves user ID from Gin context
func FromContext(c *gin.Context) string {
	user, _ := c.Get("userID")
	if userID, ok := user.(string); ok {
		return userID
	}
	return ""
}
