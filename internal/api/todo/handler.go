package api

import (
	"context"
	application "todoDB/internal/application/todo"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service *application.TodoServices
}

func NewTodoHandler(s *application.TodoServices) *todoHandler {
	return &todoHandler{service: s}
}

func (t *todoHandler) GetAllTodo(ctx context.Context, c *gin.Context) {
	t.service.GetTodoByUserName(ctx, "the user name")
}
