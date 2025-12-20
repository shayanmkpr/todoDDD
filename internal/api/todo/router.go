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

		// Task routes
		g.POST("/:todo_id/tasks", h.AddTask)
		g.PUT("/tasks", h.UpdateTask)
		g.DELETE("/tasks/:task_id", h.DeleteTask)
	}
}
