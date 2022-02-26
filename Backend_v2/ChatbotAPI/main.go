package main

import (
	"ChatbotAPI/config"
	"ChatbotAPI/hanlder"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.RedisInit()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Chat Bot API Server is running")
	})

	api := app.Group("/api")

	api.Post("/flow", hanlder.CreateFlow)
	api.Post("/action", hanlder.CreateAction)
	api.Post("/option", hanlder.CreateOption)

	app.Listen(":3010")

}
