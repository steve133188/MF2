package Routes

import (
	"mf-customer-services/Services"

	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	//contact page
	route.Get("/", Services.GetAllCustomers)
	route.Get("/id/:id", Services.GetCustomerById)
	route.Get("/name/:name", Services.GetCustomerByName)

	route.Post("/", Services.AddCustomer)
	route.Post("/addMany", Services.AddManyCustomer)
	route.Put("/id", Services.UpdateCustomerByID)
	route.Put("/many", Services.UpdateManyCustomers)
	route.Put("/phone", Services.PutPhoneToCustomer)

	route.Delete("/id", Services.DeleteCustomerById)
	route.Delete("/many", Services.DeleteManyCustomers)

	// route.Delete("/delMany", Services.DeleteManyCustomer)

	// route.Get("/chanInfo/:phone", Services.GetChannelInfoByPhone)

	route.Post("/add-tags", Services.AddTags)
	route.Put("/edit-tags", Services.UpdateCustomersTags)
	//del 1 tag from all customers
	route.Put("/del-tag", Services.DeleteTagFromAllCustomer)
	//del 1 tag from 1 customer
	route.Put("/del-customer-tag", Services.DeleteCustomerTags)
	// need to define

	//Group
	route.Get("/group/:group", Services.GetAllCustomerByGroup)
	route.Put("/add-group-to-customer", Services.AddGroupToCustomer)
	route.Put("/edit-group-name", Services.EditGroupName)
	route.Put("/detele/group/:group", Services.DeleteGrpupByName)

	//sorting
	route.Post("/filter/agent", Services.GetAgentFilter)
	route.Post("/filter/tag", Services.GetTagsFilter)
	route.Post("/filter/channel", Services.GetChannelFilter)
	// route.Post("/filter/team", Services.GetTeamFilter)

	route.Get("/team/:id", Services.GetCustomersByTeamID)
	route.Get("/team", Services.GetCustomersWithNoTeam)
	route.Put("/add-team-to-customer", Services.AddTeamIDToCustomer)
	route.Put("/change-customers-team", Services.UpdateCustomersTeamID)
	route.Put("/delete-customers-team/:team", Services.DeleteTeamIDFromCustomers)

}
