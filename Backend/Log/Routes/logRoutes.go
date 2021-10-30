package Routes

import (
	"mf-log-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/customer", Services.GetAllCustomersLog)
	route.Get("/id/:id", Services.GetCustomerLogById)
	route.Get("/type/:type", Services.GetCustomerLogByType)
	route.Get("/userid/:userid", Services.GetCustomerLogByUserId)

	route.Post("/customer", Services.CreateCustomerLog)
	route.Post("/user", Services.CreateUserLog)
	route.Post("/system", Services.CreateSystemLog)
}
