package routes

import (
	"app/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Health check
	api.Get("/health", handlers.Health)

	// Auth
	AuthRouter(api)

	// User
	UserRouter(api)
}
