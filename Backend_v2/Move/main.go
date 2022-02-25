package main

import (
	"Move/config"
	"Move/handler"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.DynamodbConfig()
	config.RedisInit()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Chatroom API Server is running")
	})

	move := app.Group("/move")
	move.Get("/chatrooms", handler.MoveChatroom)

	api := app.Group("/api")
	api.Get("/chatrooms", handler.GetChatrooms)
	api.Get("/chatrooms/user/:userId", handler.GetChatroomsByUser)

	api.Put("/chatroom/channel/:channel/room/:room_id", handler.UpdateChatroomUnreadToZero)
	app.Listen(":3010")
}
