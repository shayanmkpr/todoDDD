package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	authApi "todoDB/internal/api/middleware"
	todoApi "todoDB/internal/api/todo"
	userApi "todoDB/internal/api/user"
	todoApplication "todoDB/internal/application/todo"
	userApplication "todoDB/internal/application/user"
	user "todoDB/internal/domain/user"
	"todoDB/internal/infra/auth_jwt"
	"todoDB/internal/infra/postgres"
	"todoDB/internal/infra/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	// initializing postgres
	db, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("automigration complete")
	}

	// initializing redis
	rdb := redis.NewrdbClient()
	rdbErr := rdb.Ping(context.Background()).Err()
	if rdbErr != nil {
		log.Println(rdbErr)
		fmt.Println(rdbErr)
	}

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

	r.Use(func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		fmt.Println("=======================")
		fmt.Println("Method:", c.Request.Method)
		fmt.Println("Path:", c.Request.URL.Path)
		fmt.Println("Body:", string(bodyBytes))
		fmt.Println("=======================")

		c.Next()
	})

	// routing

	protectedGroup := r.Group("/user") // anyting going through the /user/ domain layer will be protected according to the correspoding user.
	fmt.Println("Assigning the authentication middleware ...")
	protectedGroup.Use(authHandler.AuthMiddleware())

	todoGroup := protectedGroup.Group("/todo")
	todoApi.RegisterTodoRoutes(todoGroup, todoHandler)

	unprotectedGroup := r.Group("/auth")
	userApi.RegisterUserRoutes(unprotectedGroup, userHandler)

	r.Run(":8080")
}
