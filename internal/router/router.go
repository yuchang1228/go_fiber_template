package router

import (
	"app/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Health check
	api.Get("/health", handler.Health)

	// Auth
	AuthRouter(api)

	// User
	UserRouter(api)
}
