package Routes

import (
	"mf-boardCast-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BoardCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBoardCasts)
	route.Get("/:id", Services.GetBoardCastById)
	route.Post("/", Services.AddBoardCast)
	route.Put("/:id", Services.UpdateBoardCastByID)
	route.Delete("/:id", Services.DeleteBoardCastById)
}
