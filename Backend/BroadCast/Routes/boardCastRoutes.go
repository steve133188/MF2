package Routes

import (
	"mf-broadCast-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BroadCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBroadCasts)
	route.Get("/name/", Services.GetBroadCastsByName)
	route.Get("/group/", Services.GetBroadCastsByGroup)

	route.Post("/add", Services.AddBroadCast)
	route.Post("/addMany", Services.AddManyBroadCast)

	route.Put("/name", Services.UpdateBroadCastByID)

	route.Delete("/name", Services.DeleteBroadCastByName)
}
