package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/id/:id", Services.GetUsersById)
	route.Get("/username/:username", Services.GetUserByUsername)
	route.Get("/email/:email", Services.GetUserByEmail)
	route.Get("/team/", Services.GetUserByTeam)

	route.Get("/", Services.GetAllUsers)

	route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name/:name", Services.UpdateUserByName)
	route.Put("/update-div-team", Services.UpdateDivisionAndTeam)
	route.Put("/change/status", Services.ChangeUserStatus)
	route.Put("/roles", Services.UpdateUserRole)

	route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name/:name", Services.DeleteUserByName)
}
