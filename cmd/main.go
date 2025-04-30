package main

import (
	"log"

	"app/database"
	"app/router"

	_ "app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:3000
// @BasePath /api
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})

	// app.Use(cors.New())

	database.ConnectDB()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
