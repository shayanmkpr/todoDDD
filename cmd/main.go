package main

import (
	"log"

	"github.com/gin-gonic/gin"

	userapi "yourapp/internal/api/user"
	userapp "yourapp/internal/application/user"
	userinfra "yourapp/internal/infra/postgres"
)

func main() {
	// 1. Initialize database (in infra layer)
	db, err := userinfra.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Initialize repositories (infra)
	userRepo := userinfra.NewUserRepository(db)

	// 3. Initialize application services
	userService := userapp.NewService(userRepo)

	// 4. Initialize handlers (API layer)
	userHandler := userapi.NewHandler(userService)

	// 5. Setup Gin router
	r := gin.Default()

	// 6. Register routes
	userapi.RegisterRoutes(r, userHandler)

	// 7. Start server
	r.Run(":8080")
}
