package Routes

import (
	"mf-channel-service/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllChannelInfo)
	route.Get("/id/:id", Services.GetChannelInfoById)
	// route.Get("/name/:name", Services.GetOrganizationByName)

	route.Post("/", Services.AddChannel)

	route.Put("/:id", Services.UpdateChannelById)

	route.Delete("/id/:id", Services.DeleteChannelById)
}
