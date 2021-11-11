package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	//contact page
	route.Get("/", Services.GetAllCustomers)
	route.Post("/id", Services.GetCustomerById)
	route.Post("/name", Services.GetCustomerByName)

	route.Post("/", Services.AddCustomer)
	route.Post("/addMany", Services.AddManyCustomer)
	route.Put("/id", Services.UpdateCustomerByID)
	route.Delete("/id", Services.DeleteCustomerById)
	// route.Delete("/delMany", Services.DeleteManyCustomer)

	// route.Get("/chanInfo/:phone", Services.GetChannelInfoByPhone)

	route.Post("/add-tags", Services.AddTags)
	route.Put("/edit-tags", Services.UpdateCustomersTags)
	//del 1 tag from all customers
	route.Put("/del-tag", Services.DeleteTagFromAllCustomer)
	//del 1 tag from 1 customer
	route.Put("/del-customer-tag", Services.DeleteCustomerTags)
	// need to define

	//sorting
	route.Post("/group", Services.GetAllCustomerByGroup)
	route.Post("/filter/agent", Services.GetAgentFilter)
	route.Post("/filter/tag", Services.GetTagsFilter)
	route.Post("/filter/channel", Services.GetChannelFilter)
	route.Post("/filter/team", Services.GetTeamFilter)

	route.Delete("/id", Services.DeleteCustomerById)

}
