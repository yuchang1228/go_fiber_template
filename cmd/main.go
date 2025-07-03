package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2/log"

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

	logDir := "log"
	logFile := "fiber.log"
	logPath := filepath.Join(logDir, logFile)

	_ = os.MkdirAll(logDir, os.ModePerm)

	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
		ErrorHandler:  middleware.ErrorHandler,
	})

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
