package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"

	"app/config"
	"app/pkg/i18n"

	_ "app/docs"
	"app/internal/database"
	_ "app/internal/database/migrations"
	"app/internal/middleware"
	"app/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/pressly/goose/v3"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

const logPath = "./logs/fiber.log"

var logger *zap.Logger

func setupLog() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	logger, _ = c.Build()
}

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

	setupLog()
	defer logger.Sync()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
		ErrorHandler:  middleware.ErrorHandler,
	})

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	// app.Use(cors.New())

	goose.SetLogger(goose.NopLogger())
	database.ConnectDB()
	database.Migrate()

	// i18n 初始化
	i18n.InitBundle()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config("FIBER_PORT"))))
}
