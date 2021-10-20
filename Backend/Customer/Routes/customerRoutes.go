package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllCustomers)
	route.Get("/:id", Services.GetCustomersById)
	route.Post("/", Services.AddCustomer)
	route.Put("/:id", Services.UpdateCaustomerByID)
	route.Delete("/:id", Services.DeleteCustomerById)
}
