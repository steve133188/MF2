package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	//contact page
	route.Get("/:num", Services.GetAllCustomers)
	route.Get("/id", Services.GetCustomerById)
	route.Get("/name", Services.GetCustomerByName)

	route.Post("/", Services.AddCustomer)
	route.Post("/addMany", Services.AddManyCustomer)
	route.Put("/id", Services.UpdateCustomerByID)
	route.Delete("/id", Services.DeleteCustomerById)
	// route.Delete("/delMany", Services.DeleteManyCustomer)

	// route.Get("/chanInfo/:phone", Services.GetChannelInfoByPhone)
	// route.Put("/chanInfo", Services.UpdateChannelInfoByPhone)

	route.Post("/add-tags", Services.AddTags)
	route.Put("/edit-tags", Services.UpdateCustomersTags)
	//del 1 tag from all customers
	route.Put("/del-tag", Services.DeleteTagFromAllCustomer)
	//del 1 tag from 1 customer
	route.Put("/del-customer-tag", Services.DeleteCustomerTags)
	// need to define

	//sorting
	route.Get("/group", Services.GetAllCustomerByGroup)
	route.Get("/filter/agent", Services.GetAgentFilter)
	route.Get("/filter/tag", Services.GetTagsFilter)
	route.Get("/filter/channel", Services.GetChannelFilter)
	route.Get("/filter/team", Services.GetTeamFilter)

	route.Delete("/id", Services.DeleteCustomerById)

}
