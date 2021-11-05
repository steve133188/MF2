package Routes

import (
	"mf-aoc-service/Services"

	"github.com/gofiber/fiber/v2"
)

func ChannelRoute(route fiber.Router) {
	route.Get("/", Services.GetAllChannelInfo)
	route.Get("/id/:id", Services.GetChannelInfoById)

	route.Post("/", Services.AddChannel)

	route.Put("/id/:id", Services.UpdateChannelById)

	route.Delete("/id/:id", Services.DeleteChannelById)
}

func AdminRoute(route fiber.Router) {
	route.Get("/getRole", Services.GetAllRole)
	route.Get("/getRoleById/:id", Services.GetRoleById)
	route.Get("/getRoleByName/:name", Services.GetRoleByName)
	route.Get("getAllTags", Services.GetAllTags)
	route.Get("getTags/:name", Services.GetTagsByName)

	route.Post("/addRole", Services.AddRole)
	route.Post("/addTags", Services.AddTags)

	route.Put("/putRole-nd/:id", Services.UpdateRoleByID)
	route.Put("/putRole-name/:name", Services.UpdateRoleByName)
	route.Put("/putTages/:name", Services.UpdateTagsByName)

	route.Delete("/delRole-id/:id", Services.DeleteRoleById)
	route.Delete("/delRole-name/:name", Services.DeleteRoleByName)
	route.Delete("/delTages/:name", Services.DeleteTagsByName)

	// route.Post("/addGroup", Services.AddGroup)
	// route.Put("/editGroup", Services.EditGroup)
	// route.Put("/delGroup", Services.DelGroup)
}

func OrgRoute(route fiber.Router) {
	route.Post("/add-div", Services.CreateDivision)
	route.Get("/get-div", Services.GetDivisionByName)
	route.Get("/get-alldiv", Services.GetAllDivision)
	route.Put("/edit-div", Services.UpdateDivisionByName)
	// route.Delete("/del-div", Services.DeleteDivisionByName)

	route.Put("/add-team", Services.CreateTeam)
	route.Put("/edit-team", Services.UpdateTeam)
	route.Put("/del-team", Services.DelTeam)
}
