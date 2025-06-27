package router

import (
	"app/database"
	"app/handler"
	"app/middleware"
	"app/repository"
	"app/service"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	userRepository := repository.NewUserRepository(database.GORM_DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := api.Group("/user", middleware.Protected())
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
	user.Patch("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
