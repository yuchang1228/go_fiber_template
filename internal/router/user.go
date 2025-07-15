package router

import (
	"app/internal/database"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	userRepository := repository.NewUserRepository(database.GORM_DB)
	userService := service.NewUserService(userRepository)
	userReportService := service.NewUserReportService(userRepository)
	userHandler := handler.NewUserHandler(userService, userReportService)

	user := api.Group("/user")
	user.Get("/report", userHandler.UserReport)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
	user.Patch("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
