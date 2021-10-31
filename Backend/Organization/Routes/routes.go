package Routes

import (
	"mf-organization-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllOrganization)
	// route.Get("/id/:id", Services.GetCustomersById)
	// route.Get("/name/:name", Services.GetOrganizationByName)

	route.Post("/agent", Services.AddAgent)

	route.Put("/phone/:phone", Services.UpdateOraganizationByPhone)

	route.Delete("/phone/:phone", Services.DeleteOrganizationByPhone)
}