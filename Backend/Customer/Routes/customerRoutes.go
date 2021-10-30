package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllCustomers)
	route.Get("/id/:id", Services.GetCustomersById)
	route.Get("/name/:name", Services.GetCustomersByName)

	route.Post("/", Services.AddCustomer)

	route.Put("/id/:id", Services.UpdateCustomerByID)

	route.Delete("/id/:id", Services.DeleteCustomerById)
}
