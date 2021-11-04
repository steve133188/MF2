package Routes

import (
	"mf-aoc-service/Services"

	"github.com/gofiber/fiber/v2"
)

func ChannelRoute(route fiber.Router) {
	route.Get("/", Services.GetAllChannelInfo)
	route.Get("/id/:id", Services.GetChannelInfoById)
	// route.Get("/name/:name", Services.GetOrganizationByName)

	route.Post("/", Services.AddChannel)

	route.Put("/id/:id", Services.UpdateChannelById)

	route.Delete("/id/:id", Services.DeleteChannelById)
}

func AdminRoute(route fiber.Router) {
	route.Get("/getRole", Services.GetAllRole)
	route.Get("/getRoleById/:id", Services.GetRoleById)
	route.Get("/getRoleByName/:name", Services.GetRoleByName)
	route.Get("getAllTags", Services.GetAllTags)
	route.Get("getTags/:name", Services.GetTagesByName)

	route.Post("/addRole", Services.AddRole)
	route.Post("/addTags", Services.AddTags)

	route.Put("/putRole-nd/:id", Services.UpdateRoleByID)
	route.Put("/putRole-name/:name", Services.UpdateRoleByName)
	route.Put("/putTages/:name", Services.UpdateTagsByName)

	route.Delete("/delRole-id/:id", Services.DeleteRoleById)
	route.Delete("/delRole-name/:name", Services.DeleteRoleByName)
	route.Delete("/delTages/:name", Services.DeleteTagsByName)

}

func OrgRoute(route fiber.Router) {
	// route.Get("/", Services.GetAllOrganization)
	// route.Get("/id/:id", Services.GetCustomersById)
	// route.Get("/name/:name", Services.GetOrganizationByName)
	route.Get("/get", Services.GetAllOrgInfo)
	// route.Post("/agent", Services.AddAgent)
	route.Post("/create/division", Services.CreateDivision)
	route.Post("/create/team", Services.CreateTeam)

	// route.Put("/phone/:phone", Services.UpdateOraganizationByPhone)

	route.Delete("/phone/:phone", Services.DeleteOrganizationByPhone)
}
