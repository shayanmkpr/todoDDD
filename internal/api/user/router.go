package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler) {
	fmt.Println("Registering the routs...")
	g := r.Group("/users")
	{
		g.POST("/register", h.Register)
		g.POST("/login", h.Login)
	}
}
