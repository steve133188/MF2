package Routes

import (
	"mf-aoc-service/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllChannelInfo)
	route.Get("/id/:id", Services.GetChannelInfoById)
	// route.Get("/name/:name", Services.GetOrganizationByName)

	route.Post("/", Services.AddChannel)

	route.Put("/id/:id", Services.UpdateChannelById)

	route.Delete("/id/:id", Services.DeleteChannelById)
}

func AdminRoute(route fiber.Router) {
	route.Get("/", Services.GetAllAdmins)
	route.Get("/id/:id", Services.GetAdminById)
	route.Post("/", Services.AddAdmin)
	route.Put("/id/:id", Services.UpdateAdminByID)
	route.Delete("/id/:id", Services.DeleteAdminById)
}

func OrgRoute(route fiber.Router) {
	route.Get("/", Services.GetAllOrganization)
	// route.Get("/id/:id", Services.GetCustomersById)
	// route.Get("/name/:name", Services.GetOrganizationByName)

	route.Post("/agent", Services.AddAgent)

	route.Put("/phone/:phone", Services.UpdateOraganizationByPhone)

	route.Delete("/phone/:phone", Services.DeleteOrganizationByPhone)
}
