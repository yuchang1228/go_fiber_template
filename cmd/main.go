package main

import (
	"log"
	"strconv"

	"app/database"
	"app/router"
	"app/util"

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
// @host localhost:3000
// @BasePath /api
func main() {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("../lang/active.en.toml")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.Query("lang")
		accept := c.Get("Accept-Language")

		localizer := i18n.NewLocalizer(bundle, lang, accept)

		name := c.Query("name")
		if name == "" {
			name = "Bob"
		}

		// Set title message.
		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})

		// Parse and set unread count of emails.
		unreadEmailCount, _ := strconv.ParseInt(c.Query("unread"), 10, 64)

		// Set your own message for unread emails.
		myUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "MyUnreadEmails",
				Description: "The number of unread emails I have",
				One:         "I have {{.PluralCount}} unread email.",
				Other:       "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
		})

		// Set other personal message for unread emails.
		personUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "PersonUnreadEmails",
				Description: "The number of unread emails a person has",
				One:         "{{.Name}} has {{.UnreadEmailCount}} unread email.",
				Other:       "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
			TemplateData: map[string]interface{}{
				"Name":             name,
				"UnreadEmailCount": unreadEmailCount,
			},
		})

		// Return rendered template.
		return c.Render("index", fiber.Map{
			"Title": helloPerson,
			"Paragraphs": []string{
				myUnreadEmails,
				personUnreadEmails,
			},
		})
	})

	// app.Use(cors.New())

	database.ConnectDB()
	util.ValidateStruct()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
