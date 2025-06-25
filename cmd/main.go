package main

import (
	"fmt"
	"log"

	"app/database"
	"app/middleware"
	"app/router"

	_ "app/docs"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:9000
// @BasePath /api
func main() {
	bundle := i18n.NewBundle(language.TraditionalChinese)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("lang/active.zh_tw.toml")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
		ErrorHandler:  middleware.ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.Query("lang")
		accept := c.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		// 簡單翻譯
		hello, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "HelloWorld",
		})
		fmt.Println(hello) // 輸出：哈囉，世界

		// 變數替換
		greeting, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "GreetingName",
			TemplateData: map[string]string{
				"Name": "小明",
			},
		})

		fmt.Println(greeting) // 輸出：哈囉，小明

		return c.SendString(hello)
	})

	// app.Use(cors.New())

	database.ConnectDB()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
