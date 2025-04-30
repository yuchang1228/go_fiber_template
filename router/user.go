package router

import (
	"app/handler"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
}
