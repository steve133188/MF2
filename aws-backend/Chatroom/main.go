package main

import (
	"aws-lambda-chatroom/config"
	"aws-lambda-chatroom/handler"
	"net/http"

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
		return c.Status(http.StatusOK).SendString("Chatroom API Server is running")
	})

	api := app.Group("/api")
	api.Get("/chatrooms", handler.GetChatrooms)
	api.Get("/chatroom/channel/:channel/room/:room_id", handler.GetOneChatroom)
	api.Get("/chatrooms/user/:userId", handler.GetChatroomsByUser)

	api.Put("/chatroom/channel/:channel/room/:room_id", handler.UpdateChatroomunreadToZero)

	api.Post("/chatrooms/user", handler.GetChatroomByAgent)

	app.Listen(":3010")
}
