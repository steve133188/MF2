package main

import (
	"mf-bot-services/DB"
	"mf-bot-services/Routes"
	"mf-bot-services/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-BotMessages-Services"})
	})

	api := app.Group("/api")
	Routes.BotBuildRoute(api.Group("/botmessages"))

	//listen to whtasapp response
	go func() {
		http.HandleFunc("/webhook", Services.HandleWhatsappResponse)
		http.ListenAndServe(":3005", nil)
	}()

	app.Listen(":3005")

}
