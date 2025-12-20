package api

import (
	"net/http"
	authMiddleware "todoDB/internal/api/middleware"
	application "todoDB/internal/application/todo"

	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	UserName string
	Title    string
}

type todoHandler struct {
	service *application.TodoServices
}

func NewTodoHandler(s *application.TodoServices) *todoHandler {
	return &todoHandler{service: s}
}

func (t *todoHandler) GetUserTodo(c *gin.Context) {
	userName := authMiddleware.UserNameFromContext(c)
	theTodo, err := t.service.GetTodoByUserName(c.Request.Context(), userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, theTodo)
}

func (t *todoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.service.CreateTodo(c.Request.Context(), req.UserName, req.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "todo created"})
}
