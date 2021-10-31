package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/id/:id", Services.GetUsersById)
	route.Get("/username/:username", Services.GetUserByUsername)
	route.Get("/email/:email", Services.GetUserByEmail)
	route.Get("/team/:team", Services.GetUserByTeam)

	route.Get("/", Services.GetAllUsers)

	route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name/:name", Services.UpdateUserByName)
	route.Put("/division/:div/team/:team/name/:name", Services.UpdateDivisionAndTeam)

	route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name/:name", Services.DeleteUserByName)
}
