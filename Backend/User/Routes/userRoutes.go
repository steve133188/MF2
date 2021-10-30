package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/id/:id", Services.GetUsersById)
	route.Get("/username/:username", Services.GetUserByUsername)
	route.Get("/email/:email", Services.GetUserByEmail)

	route.Get("/", Services.GetAllUsers)

	route.Put("/id/:id", Services.UpdateUserByID)
	route.Delete("/id/:id", Services.DeleteUserById)
}
