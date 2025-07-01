package router

import (
	"app/internal/database"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	userRepository := repository.NewUserRepository(database.GORM_DB)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
}
