package Routes

import (
	"mf-chat-services/Services"

	"github.com/gofiber/fiber/v2"
)

func ChatRoute(route fiber.Router) {
	route.Get("/", Services.GetAllMessages)
	route.Get("/id/:id", Services.GetOneMessageById)
	route.Post("/", Services.AddOneMessage)
	route.Put("/id/:id", Services.UpdateOneMessageById)
	route.Delete("/id/:id", Services.DeleteOneMessageById)
}
