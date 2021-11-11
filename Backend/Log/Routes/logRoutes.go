package Routes

import (
	"mf-log-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Get("/customer/name", Services.GetAllCustomersLogByName)
	route.Get("/userid/:userid", Services.GetCustomerLogByUserId)

	route.Post("/customer", Services.CreateCustomerLog)
	route.Post("/user", Services.CreateUserLog)
	route.Post("/system", Services.CreateSystemLog)

	route.Post("/manyUser", Services.AddManyUserLog)
	route.Get("/manyUser", Services.GetManyUserLog)
}
