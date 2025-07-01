package main

import (
	"fmt"
	"log"

	"app/config"
	"app/pkg/i18n"

	_ "app/docs"
	"app/internal/database"
	_ "app/internal/database/migrations"
	"app/internal/job"
	"app/internal/middleware"
	"app/internal/router"

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

	job.InitRabbitMQ()
	defer job.Conn.Close()
	defer job.Channel.Close()

	// i18n 初始化
	i18n.InitBundle()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config("FIBER_PORT"))))
}
