package Routes

import (
	"mf-message-services/Services"

	"github.com/gofiber/fiber/v2"
)

func MessageRoutes(route fiber.Router) {
	route.Get("/user/:id", Services.GetByUserID)
	route.Get("/chatroom/:id", Services.GetByChatroomID)

}
