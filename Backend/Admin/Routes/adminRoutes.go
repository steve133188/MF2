package Routes

import (
	"mf-admin-services/Services"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(route fiber.Router) {
	route.Get("/", Services.GetAllAdmins)
	route.Get("/:id", Services.GetAdminById)
	route.Post("/", Services.AddAdmin)
	route.Put("/:id", Services.UpdateAdminByID)
	route.Delete("/:id", Services.DeleteAdminById)
}
