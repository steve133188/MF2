package Routes

import (
	"mf-log-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Post("/customer", Services.CreateCustomerLog)
	route.Post("/user", Services.CreateUserLog)
	route.Post("/system", Services.CreateSystemLog)
}
