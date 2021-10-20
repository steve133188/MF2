package main

import (
	"mf-chat-services/DB"
	"mf-chat-services/Routes"

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
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Chats-Services"})
	})

	api := app.Group("/api")

	Routes.ChatRoute(api.Group("/messages"))

	app.Listen(":3003")
}
