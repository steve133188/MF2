package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/:id", Services.GetUsersById)
	route.Get("/:username", Services.GetUserByUsername)
	route.Get("/", Services.GetAllUsers)

	route.Post("/", Services.AddUser)
	route.Post("/login", Services.Login)

	route.Put("/:id", Services.UpdateUserByID)
	route.Delete("/:id", Services.DeleteUserById)
}
