package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	//contact page
	route.Get("/all", Services.GetAllCustomers)
	route.Get("/id", Services.GetCustomerById)
	route.Get("/name", Services.GetCustomerByName)
	route.Post("/add", Services.AddCustomer)
	route.Post("/addMany", Services.AddManyCustomer)
	route.Put("/id", Services.UpdateCustomerByID)
	route.Delete("/id", Services.DeleteCustomerById)
	// route.Delete("/delMany", Services.DeleteManyCustomer)

	// route.Get("/chanInfo/:phone", Services.GetChannelInfoByPhone)
	// route.Put("/chanInfo", Services.UpdateChannelInfoByPhone)

	route.Post("/addTags", Services.AddTags)
	// need to define
	// route.Put("/editTags", Services.UpdateCustomerTags)

	//sorting
	route.Get("/group", Services.GetAllCustomerByGroup)
	route.Get("/filter/agent", Services.GetAgentFilter)
	route.Get("/filter/tag", Services.GetTagsFilter)
	route.Get("/filter/channel", Services.GetChannelFilter)
	// route.Get("/filter/group", Services.GetGroupFilter)

	route.Delete("/phone", Services.DeleteCustomerByPhone)
}
