package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	g := r.Group("/users")
	{
		g.POST("/register", h.Register)
		g.POST("/login", h.Login)
	}
}
