package Routes

import (
	"mf-user-servies/Services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(route fiber.Router) {
	route.Get("/name/:name", Services.GetUserByName)
	route.Get("/email/:email", Services.GetUserByEmail)
	route.Get("/phone/:phone", Services.GetUserByPhone)
	route.Get("/", Services.GetAllUsers)
	route.Get("/userlist", Services.GetUserList)

	// route.Get("/authority", Services.GetUserAuthority)
	//for filter

	// route.Put("/id/:id", Services.UpdateUserByID)
	route.Put("/name", Services.UpdateUserByName)
	route.Put("/status", Services.ChangeUserStatus)

	//role

	//team
	route.Get("/team/:id", Services.GetUsersByTeamID)
	route.Put("/add-team-to-user", Services.AddTeamIDToUser)
	route.Put("/change-users-team", Services.UpdateUsersTeamID)
	route.Put("/delete-users-team/:team", Services.DeleteTeamIDFromUsers)

	// route.Delete("/id/:id", Services.DeleteUserById)
	route.Delete("/name/:name", Services.DeleteUserByName)

	route.Put("/chanInfo/:phone", Services.UpdateChannelInfoByPhone)

	route.Get("/role-number/:role", Services.GetUserNumByRole)
	route.Get("/role-auth/:phone", Services.GetUserAuthByPhone)
	route.Get("/role/:role", Services.GetUsersByRole)
	route.Put("/role", Services.UpdateUserRole)
	route.Put("/roles", Services.UpdateUsersRole)
	route.Put("/role/:phone", Services.DeleteUserRoleByPhone)
	route.Put("/roles/:role", Services.DeleteUsersRoleByRole)
	route.Post("/role", Services.AddRoleToUser)

}
