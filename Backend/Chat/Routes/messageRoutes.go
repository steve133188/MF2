package Routes

import (
	"mf-chat-services/Services"

	"github.com/gofiber/fiber/v2"
)

func ChatRoute(route fiber.Router) {
	route.Get("/getAll", Services.GetAllMessages)

	route.Post("/postOne", Services.AddOneMessage)
	route.Post("/postMany", Services.AddManyMessages)

}
