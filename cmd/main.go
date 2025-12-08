package main

import (
	"log"

	api "todoDB/internal/api/user"
	application "todoDB/internal/application/user"
	"todoDB/internal/infra/auth_jwt"
	"todoDB/internal/infra/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	// get the db as an object
	db, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := postgres.NewUserRepository(db)
	authRepo := auth_jwt.NewAuthRepository()
	userService := application.NewService(userRepo, authRepo)
	userHandler := api.NewHandler(userService)
	r := gin.Default()
	api.RegisterRoutes(r, userHandler)
	r.Run(":8080")
}
