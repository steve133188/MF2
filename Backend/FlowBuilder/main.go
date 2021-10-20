package main

import (
	"mf-flowbuilder-services/DB"
	"mf-flowbuilder-services/Routes"
	"mf-flowbuilder-services/Services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	// app.Get("/test", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-BOtBuilds-Services"})
	// })

	app.Post("/", Services.Testing)
	api := app.Group("/api")

	Routes.BotBuildRoute(api.Group("/bots"))

	app.Listen(":3006")

}
