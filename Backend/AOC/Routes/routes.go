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
	route.Get("/tags", Services.GetAllTags)
	route.Get("/tag/name/:name", Services.GetTagByName)
	route.Get("/taglist", Services.GetTagList)
	route.Post("/tags", Services.AddTag)
	route.Put("/tage/name/:name", Services.UpdateTagsByName)
	route.Delete("/tages/name/:name", Services.DeleteTagsByName)

	//role CRUD
	route.Get("/roles", Services.GetAllRoles)
	route.Get("/role/name/:name", Services.GetRoleByName)
	route.Get("/roles-name", Services.GetRolesName)
	route.Post("/role", Services.AddRole)
	route.Put("/role", Services.UpdateRoleByName)
	route.Delete("/role/name/:name", Services.DeleteRoleByName)
	// route.Post("/addGroup", Services.AddGroup)
	// route.Put("/editGroup", Services.EditGroup)
	// route.Put("/delGroup", Services.DelGroup)

	//Standard Reply
	route.Get("/reply/id/:id", Services.GetReplyFolderByID)
	route.Get("/replys", Services.GetAllReplyFolder)
	route.Post("/reply", Services.CreateReply)
	route.Put("/add-content", Services.AddContent)
	route.Put("/edit-content", Services.UpdateContent)
	route.Put("/del-content", Services.DeleteContent)
	route.Get("/content/:name", Services.GetContentsByFolderName)

	route.Put("/reply", Services.UpdateReply)
	route.Delete("/reply/:id", Services.DeleteReply)
}

func OrgRoute(route fiber.Router) {
	route.Post("/", Services.CreateDivision)
	route.Get("/root", Services.GetRootDivisions)
	route.Get("/parent/:parentId", Services.GetOrgByParentID)
	route.Get("/id/:id", Services.GetOrgByID)
	route.Get("/name/:id", Services.GetNameByID)

	route.Put("/", Services.EditOrgName)
	route.Delete("/id/:id", Services.DeleteOrgById)
}
