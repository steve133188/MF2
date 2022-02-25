package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf2-channel-router/handler"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// redis initialization
	//config.RedisInit()

	api := app.Group("/ch-router")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("MF2 Channel Router is running")
	})

	api.Post("/connect", handler.ChannelConnect)
	api.Get("/restart", handler.ChannelRestart)
	api.Get("/disconnect", handler.ChannelDisconnect)
	api.Post("/send-message", handler.ChannelSendMessage)

	api.Get("/update-status", handler.ChannelUpdateStatus)

	app.Listen(":8081")
}
