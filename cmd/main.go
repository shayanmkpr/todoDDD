package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	authApi "todoDB/internal/api/middleware"
	userApi "todoDB/internal/api/user"
	application "todoDB/internal/application/user"
	user "todoDB/internal/domain/user"
	"todoDB/internal/infra/auth_jwt"
	"todoDB/internal/infra/postgres"
	"todoDB/internal/infra/redis"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("automigration complete")
	}

	rdb := redis.NewrdbClient()
	rdbErr := rdb.Ping(context.Background()).Err()
	if rdbErr != nil {
		log.Println(rdbErr)
		fmt.Println(rdbErr)
	}

	userRepo := postgres.NewUserRepository(db)
	authRepo := auth_jwt.NewAuthRepository()
	refRepo := redis.NewRdbRepository(rdb)

	userService := application.NewService(userRepo, authRepo, refRepo)
	userHandler := userApi.NewHandler(userService)
	authHandler := authApi.NewAuthHandler(userService)

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

	protectedGroup := r.Group()

	authHandler.AuthMiddleware()
	userApi.RegisterRoutes(r, userHandler)

	r.Run(":8080")
}
