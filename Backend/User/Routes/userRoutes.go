package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/name/:name", Services.GetUserByName)
	route.Get("/email/:email", Services.GetUserByEmail)
	route.Get("/phone/:phone", Services.GetUserByPhone)
	route.Get("/team", Services.GetUsersByTeam)
	route.Get("/", Services.GetAllUsers)
	route.Get("/userlist", Services.GetUserList)
	// route.Get("/authority", Services.GetUserAuthority)
	//for filter

	// route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name/:name", Services.UpdateUserByName)
	route.Put("/update-div-team", Services.UpdateDivisionAndTeam)
	route.Put("/status", Services.ChangeUserStatus)
	route.Put("/role", Services.UpdateUserRole) //??????
	route.Put("/team", Services.DeleteUserTeam)

	// route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name/:name", Services.DeleteUserByName)

	route.Put("/chanInfo/:phone", Services.UpdateChannelInfoByPhone)

}

func RoleRoute(route fiber.Router) {
	route.Get("/", Services.GetAllRoles)
	route.Get("/name/:name", Services.GetRoleByName)
	route.Get("/list", Services.GetRolesName)
	route.Post("/", Services.AddRole)
	route.Put("/", Services.UpdateRoleByName)
	route.Delete("/name/:name", Services.DeleteRoleByName)
}
