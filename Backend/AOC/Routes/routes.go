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

	route.Post("/addRole", Services.AddRole)

	route.Put("/putRole-nd/:id", Services.UpdateRoleByID)
	route.Put("/putRole-name/:name", Services.UpdateRoleByName)

	route.Delete("/delRole-id/:id", Services.DeleteRoleById)
	route.Delete("/delRole-name/:name", Services.DeleteRoleByName)

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
