package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	// route.Get("/id/:id", Services.GetUsersById)
	route.Get("/username", Services.GetUserByUsername)
	route.Get("/email", Services.GetUserByEmail)

	route.Get("/phone", Services.GetUserByPhone)
	route.Get("/team", Services.GetUsersByTeam)
	// route.Get("/authority", Services.GetUserAuthority)
	//for filter
	route.Get("/userlist", Services.GetUserList)
	route.Get("/", Services.GetAllUsers)

	// route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name/", Services.UpdateUserByName)
	route.Put("/update-div-team", Services.UpdateDivisionAndTeam)
	route.Put("/change/status", Services.ChangeUserStatus)
	route.Put("/roles", Services.UpdateUserRole)
	route.Put("/team", Services.DeleteUserTeam)

	// route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name", Services.DeleteUserByName)

}
