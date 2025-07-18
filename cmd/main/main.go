package main

import (
	"fmt"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"app/config"
	"app/pkg/i18n"
	"app/util"

	_ "app/docs"
	"app/internal/databases"
	_ "app/internal/databases/migrations"
	_ "app/internal/handlers"
	"app/internal/middlewares"
	"app/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/pressly/goose/v3"
	// "github.com/gofiber/fiber/v2/middlewares/cors"
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
	loc, _ := time.LoadLocation(config.Config("TIMEZONE"))
	time.Local = loc

	util.SetupLog()
	defer util.Logger.Sync()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
		ErrorHandler:  middlewares.ErrorHandler,
	})

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: util.Logger,
	}))

	app.Use(recover.New())

	// app.Use(cors.New())

	goose.SetLogger(goose.NopLogger())
	config.ConnectDB()
	databases.Migrate()

	// i18n 初始化
	i18n.InitBundle()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config("FIBER_PORT"))))
}
