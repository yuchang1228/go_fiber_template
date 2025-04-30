package router

import (
	"app/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
}
