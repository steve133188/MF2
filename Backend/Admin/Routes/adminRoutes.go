package Routes

import (
	"mf-admin-services/Services"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(route fiber.Router) {
	route.Get("/", Services.GetAllAdmins)
	route.Get("/id/:id", Services.GetAdminById)
	route.Post("/", Services.AddAdmin)
	route.Put("/id/:id", Services.UpdateAdminByID)
	route.Delete("/id/:id", Services.DeleteAdminById)
}
