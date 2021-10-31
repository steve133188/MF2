package Routes

import (
	"mf-boardCast-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BoardCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBoardCasts)
	route.Get("/id/:id", Services.GetBoardCastById)
	route.Get("/name/:name", Services.GetBoardCastByName)
	route.Get("/group/:group", Services.GetBoardCastsByGroup)

	route.Post("/", Services.AddBoardCast)

	route.Put("/id/:id", Services.UpdateBoardCastByID)

	route.Delete("/id/:id", Services.DeleteBoardCastById)
	route.Delete("/name/:name", Services.DeleteBoardCastByName)
}
