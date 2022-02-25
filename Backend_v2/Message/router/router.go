package router

import (
	"github.com/gofiber/fiber/v2"
	"mf2-message-server/handler"
)

// message
func MsgRouter(router fiber.Router) {
	router.Post("/", handler.AddMessage)
}

// messages
func MsgsRouter(router fiber.Router) {
	router.Get("/chatroom/:roomId", handler.GetAllMessagesByChatroom)
}
