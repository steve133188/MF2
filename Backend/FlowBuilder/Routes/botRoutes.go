package Routes

import (
	"mf-flowbuilder-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BotBuildRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBots)
	route.Get("/:id", Services.GetOneBotById)
	route.Post("/", Services.CreateOneBot)
	route.Put("/:id", Services.UpdateOneBotById)
	route.Delete("/:id", Services.DeleteOneBotById)

	// route.Post("/webhook-newbody", Services.NewReqBody)
}
