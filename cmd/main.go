package main

import (
	"fmt"
	"log"

	"app/config"
	"app/database"
	"app/middleware"
	"app/router"
	"app/util"

	_ "app/database/migrations"
	_ "app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:9000
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
		ErrorHandler:  middleware.ErrorHandler,
	})

	// app.Use(cors.New())

	database.ConnectDB()
	database.Migrate()

	// i18n 初始化
	util.InitBundle()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config("FIBER_PORT"))))
}
