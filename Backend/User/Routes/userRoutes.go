package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	// route.Get("/id/:id", Services.GetUsersById)
	route.Post("/username", Services.GetUserByUsername)
	route.Post("/email", Services.GetUserByEmail)

	route.Post("/phone", Services.GetUserByPhone)
	route.Post("/team", Services.GetUsersByTeam)
	// route.Get("/authority", Services.GetUserAuthority)
	//for filter
	route.Post("/userlist", Services.GetUserList)
	route.Post("/", Services.GetAllUsers)

	// route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name/", Services.UpdateUserByName)
	route.Put("/update-div-team", Services.UpdateDivisionAndTeam)
	route.Put("/change/status", Services.ChangeUserStatus)
	route.Put("/roles", Services.UpdateUserRole)
	route.Put("/team", Services.DeleteUserTeam)

	// route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name", Services.DeleteUserByName)

	route.Put("/chanInfo", Services.UpdateChannelInfoByID)

}
