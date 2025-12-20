package api

import (
	"context"
	"net/http"

	application "todoDB/internal/application/user"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *application.UserService
}

func NewAuthHandler(s *application.UserService) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

func (a *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		authClaims, err := a.authService.ValidateAccessToken(context.Background(), token) // not sure about this
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// attach user info to context
		c.Set("userName", authClaims.UserName)
		c.Next()
	}
}

func FromContext(c *gin.Context) string {
	user, _ := c.Get("userName")
	if userID, ok := user.(string); ok {
		return userID
	}
	return ""
}
