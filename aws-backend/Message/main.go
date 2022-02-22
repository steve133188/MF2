package main

import (
	"aws-message-api/config"
	"aws-message-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.DynamodbConfig()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Message API Server is running")
	})

	api := app.Group("/api")
	api.Get("/messages/chatroom/:roomId", handler.GetAllMessagesByChatroom)
	api.Post("/message", handler.AddMessage)

	app.Listen(":3011")
}
