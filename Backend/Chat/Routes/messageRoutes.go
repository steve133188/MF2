package Routes

import (
	"mf-chat-services/Services"

	"github.com/gofiber/fiber/v2"
)

func ChatRoute(route fiber.Router) {
	route.Get("/", Services.GetAllMessages)
	route.Get("/:id", Services.GetOneMessageById)
	route.Post("/", Services.AddOneMessage)
	route.Put("/:id", Services.UpdateOneMessageById)
	route.Delete("/:id", Services.DeleteOneMessageById)
}
