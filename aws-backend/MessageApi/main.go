package main

import (
	"mf2-message-api/config"
	"mf2-message-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.DynamodbConfig()

	api := app.Group("/msg-api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Message API Server is running")
	})

	router.MsgRouter(api.Group("/message"))
	router.MsgsRouter(api.Group("/messages"))
	router.CustomerRouter(api.Group("/customer"))
	router.ChatRouter(api.Group("/chatroom"))
	router.ChatsRouter(api.Group("/chatrooms"))
	router.ActivityRouter(api.Group("/activity"))

	app.Listen(":8080")
}
