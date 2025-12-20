package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(g *gin.RouterGroup, h *todoHandler) {
	fmt.Println("Todo Routes")
	{
		g.GET("/", h.GetUserTodo)
		g.POST("/", h.CreateTodo)
	}
}
