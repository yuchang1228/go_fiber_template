package router

import (
	"app/database"
	"app/handler"
	"app/repository"
	"app/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	userRepository := repository.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
}
