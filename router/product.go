package router

import (
	"app/handler"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(api fiber.Router) {
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
