package Routes

import (
	"mf-boardCast-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BoardCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBoardCasts)
	route.Get("/name/", Services.GetBoardCastsByName)
	route.Get("/group/", Services.GetBoardCastsByGroup)

	route.Post("/add", Services.AddBoardCast)
	route.Post("/addMany", Services.AddManyBoardCast)

	route.Put("/name", Services.UpdateBoardCastByID)

	route.Delete("/name", Services.DeleteBoardCastByName)
}
