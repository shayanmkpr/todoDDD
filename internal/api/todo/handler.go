package api

import (
	"net/http"
	"strconv"

	authMiddleware "todoDB/internal/api/middleware"
	application "todoDB/internal/application/todo"
	"todoDB/internal/domain/todo"

	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required"`
	Status string `json:"status"`
	Body   string `json:"body"`
}

type UpdateTaskRequest struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Body   string `json:"body"`
}

type todoHandler struct {
	service *application.TodoServices
}

func NewTodoHandler(s *application.TodoServices) *todoHandler {
	return &todoHandler{service: s}
}

func (t *todoHandler) GetUserTodo(c *gin.Context) {
	userName := authMiddleware.UserNameFromContext(c)
	todos, err := t.service.GetTodoByUserName(c.Request.Context(), userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (t *todoHandler) CreateTodo(c *gin.Context) {
	userName := authMiddleware.UserNameFromContext(c)

	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.service.CreateTodo(c.Request.Context(), userName, req.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "todo created"})
}

func (t *todoHandler) AddTask(c *gin.Context) {
	todoIDParam := c.Param("todo_id")
	todoID, err := strconv.Atoi(todoIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo ID"})
		return
	}

	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &todo.Task{
		Title:  req.Title,
		Status: req.Status,
		Body:   req.Body,
	}

	if err := t.service.AddTask(c.Request.Context(), todoID, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task added"})
}

func (t *todoHandler) UpdateTask(c *gin.Context) {
	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &todo.Task{
		ID:     req.ID,
		Title:  req.Title,
		Status: req.Status,
		Body:   req.Body,
	}

	if err := t.service.UpdateTask(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func (t *todoHandler) DeleteTask(c *gin.Context) {
	taskIDParam := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := t.service.DeleteTask(c.Request.Context(), taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
