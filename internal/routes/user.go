package routes

import (
	"app/config"
	"app/internal/handlers"
	"app/internal/repositories"
	"app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	userRepository := repositories.NewUserRepository(config.GORM_DB)
	userService := services.NewUserService(userRepository)
	userReportService := services.NewUserReportService(userRepository)
	userHandler := handlers.NewUserHandler(userService, userReportService)

	user := api.Group("/user")
	user.Get("/report", userHandler.UserReport)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
	user.Patch("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
