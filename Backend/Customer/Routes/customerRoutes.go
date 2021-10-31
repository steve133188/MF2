package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllCustomers)
	route.Get("/id/:id", Services.GetCustomersById)
	route.Get("/name/:name", Services.GetCustomersByName)
	route.Get("/team/:team", Services.GetAllByTeamSorting)

	route.Get("/sort/agent", Services.GetAgentSorting)
	route.Get("/sort/tag", Services.GetTagsSorting)
	route.Get("/sort/channel", Services.GetChannelSorting)

	route.Post("/create", Services.AddCustomer)

	route.Put("/id/:id", Services.UpdateCustomerByID)

	route.Delete("/id/:id", Services.DeleteCustomerById)
	route.Delete("/delete", Services.DeleteCustomer)
}
