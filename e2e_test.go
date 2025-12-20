package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	authApi "todoDB/internal/api/middleware"
	todoApi "todoDB/internal/api/todo"
	userApi "todoDB/internal/api/user"
	todoApplication "todoDB/internal/application/todo"
	userApplication "todoDB/internal/application/user"
	userDomain "todoDB/internal/domain/user"
	todoDomain "todoDB/internal/domain/todo"
	"todoDB/internal/infra/auth_jwt"
	"todoDB/internal/infra/postgres"
	"todoDB/internal/infra/redis"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestPostgres() (*gorm.DB, error) {
	// Using SQLite in-memory for testing
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}

func setupTestServer() *gin.Engine {
	// Using in-memory SQLite for testing
	db, err := NewTestPostgres()
	if err != nil {
		// If test DB fails, we'll try the regular one
		db, err = postgres.NewPostgres()
		if err != nil {
			panic(err)
		}
	}

	// Clean up and recreate tables for tests
	db.Migrator().DropTable(&userDomain.User{})
	db.Migrator().DropTable(&todoDomain.Todo{})
	db.Migrator().DropTable(&todoDomain.Task{})

	db.AutoMigrate(&userDomain.User{}, &todoDomain.Todo{}, &todoDomain.Task{})

	rdb := redis.NewrdbClient()
	userRepo := postgres.NewUserRepository(db)
	todoRepo := postgres.NewTodoRepository(db)
	authRepo := auth_jwt.NewAuthRepository()
	refRepo := redis.NewRdbRepository(rdb)

	userService := userApplication.NewService(userRepo, authRepo, refRepo)
	todoService := todoApplication.NewTodoService(todoRepo, authRepo)

	userHandler := userApi.NewHandler(userService)
	authHandler := authApi.NewAuthHandler(userService)
	todoHandler := todoApi.NewTodoHandler(todoService)

	r := gin.Default()

	protectedGroup := r.Group("/user")
	protectedGroup.Use(authHandler.AuthMiddleware())

	todoGroup := protectedGroup.Group("/todo")
	todoApi.RegisterTodoRoutes(todoGroup, todoHandler)

	unprotectedGroup := r.Group("/auth")
	userApi.RegisterUserRoutes(unprotectedGroup, userHandler)

	return r
}

func TestEndToEndFlow(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestServer()

	// Test User Registration
	userName := fmt.Sprintf("test_user_%d", time.Now().Unix())
	password := "test_password_123"

	registerReq := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{
		UserName: userName,
		Password: password,
	}

	registerBody, _ := json.Marshal(registerReq)
	registerResp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(registerBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(registerResp, req)

	assert.Equal(t, http.StatusCreated, registerResp.Code, "Expected registration to succeed")
	fmt.Printf("Register Response Status: %d\n", registerResp.Code)

	// Test Login
	loginReq := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{
		UserName: userName,
		Password: password,
	}

	loginBody, _ := json.Marshal(loginReq)
	loginResp := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(loginResp, req)

	assert.Equal(t, http.StatusOK, loginResp.Code, "Expected login to succeed")

	var loginResponse map[string]interface{}
	json.Unmarshal(loginResp.Body.Bytes(), &loginResponse)
	accessToken, ok := loginResponse["access_token"].(string)
	if !ok {
		// Try the key as access_token
		t.Errorf("Could not extract access token from response: %v", loginResponse)
		return
	}
	assert.NotEmpty(t, accessToken, "Access token should not be empty")

	fmt.Printf("Login Response Status: %d\n", loginResp.Code)

	// Test Create Todo with Authorization
	todoTitle := "Test Todo"
	todoReq := todoApi.CreateTodoRequest{
		Title: todoTitle,
	}

	todoBody, _ := json.Marshal(todoReq)
	todoResp := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/user/todo/", bytes.NewBuffer(todoBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)
	r.ServeHTTP(todoResp, req)

	assert.Equal(t, http.StatusCreated, todoResp.Code, "Expected todo creation to succeed")
	fmt.Printf("Create Todo Response Status: %d\n", todoResp.Code)

	// Test Get User Todos
	getTodosResp := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/todo/", nil)
	req.Header.Set("Authorization", accessToken)
	r.ServeHTTP(getTodosResp, req)

	assert.Equal(t, http.StatusOK, getTodosResp.Code, "Expected to get todos successfully")
	fmt.Printf("Get Todos Response Status: %d\n", getTodosResp.Code)
	fmt.Printf("Get Todos Response Body: %s\n", getTodosResp.Body.String())

	// Try to parse the todos to get an actual todo ID
	var todos []todoDomain.Todo
	err := json.Unmarshal(getTodosResp.Body.Bytes(), &todos)
	if err != nil {
		t.Errorf("Could not parse todos response: %v", err)
	} else if len(todos) > 0 {
		// Use the first todo ID to add a task
		todoID := todos[0].ID
		taskTitle := "Test Task"
		taskReq := struct {
			Title  string `json:"title"`
			Status string `json:"status"`
			Body   string `json:"body"`
		}{
			Title:  taskTitle,
			Status: "pending",
			Body:   "Task body",
		}

		taskBody, _ := json.Marshal(taskReq)
		taskResp := httptest.NewRecorder()
		req, _ = http.NewRequest("POST", fmt.Sprintf("/user/todo/%d/tasks", todoID), bytes.NewBuffer(taskBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", accessToken)
		r.ServeHTTP(taskResp, req)

		assert.Equal(t, http.StatusCreated, taskResp.Code, "Expected task creation to succeed")
		fmt.Printf("Add Task to Actual Todo Response Status: %d\n", taskResp.Code)
		fmt.Printf("Add Task Response Body: %s\n", taskResp.Body.String())
	}

	fmt.Println("End-to-end test completed successfully!")
}