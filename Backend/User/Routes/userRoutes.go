package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/", Services.GetAllUsers)
	route.Get("/:id", Services.GetUsersById)
	route.Get("/:username", Services.GetUserByUsername)
	route.Post("/", Services.AddUser)
	route.Put("/:id", Services.UpdateUserByID)
	route.Delete("/:id", Services.DeleteUserById)
}
