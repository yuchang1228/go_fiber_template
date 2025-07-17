package routes

import (
	"app/config"
	"app/internal/handlers"
	"app/internal/repositories"
	"app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	userRepository := repositories.NewUserRepository(config.GORM_DB)
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
}
