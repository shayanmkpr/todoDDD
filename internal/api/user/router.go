package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(g *gin.RouterGroup, h *Handler) {
	fmt.Println("User Routes")
	{
		g.POST("/register", h.Register)
		g.POST("/login", h.Login)
		g.POST("/ref", h.ValidateRefreshToken)
	}
}
