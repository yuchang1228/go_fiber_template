package router

import (
	"app/handler"
	"app/repository"
	"app/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService)

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
}
