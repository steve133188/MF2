package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf2-broadcast/handler"
)

func main() {
	app := fiber.New()

	//config.DynamodbConfig()

	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api-v2")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("BroadCast Module is running")
	})

	api.Post("/broadcast", handler.SendBCTest)
	api.Post("/test/message", handler.AddMessageTest)
	api.Post("/template", handler.CreateWtsTemplate)

	app.Listen(":8082")
}
