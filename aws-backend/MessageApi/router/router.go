package router

import (
	"mf2-message-api/handler"

	"github.com/gofiber/fiber/v2"
)

// /message
func MsgRouter(router fiber.Router) {
	router.Post("/", handler.AddMessage)
}

// /messages
func MsgsRouter(router fiber.Router) {
	router.Get("/chatroom/:roomId", handler.GetAllMessagesByChatroom)
}

// /chatroom
func ChatRouter(router fiber.Router) {
	router.Put("/channel/:channel/room/:room_id", handler.UpdateChatroomunreadToZero)
	router.Get("/channel/:channel/room/:room_id", handler.CheckChatroom)
	router.Post("/", handler.AddChatroom)
	router.Put("/name", handler.UpdateChatroomName)
}

// /chatrooms
func ChatsRouter(router fiber.Router) {
	router.Post("/user", handler.GetChatroomsByAgent)
}

// /customer
func CustomerRouter(router fiber.Router) {
	router.Post("/", handler.AddCustomer)
	router.Get("/:id", handler.GetCustomer)
}

// /activity
func ActivityRouter(router fiber.Router) {
	router.Post("/", handler.AddActivity)
}
