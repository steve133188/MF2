package Routes

import (
	"mf-bot-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BotBuildRoute(route fiber.Router) {
	route.Get("/", Services.GetAllBots)
	route.Get("/id/:id", Services.GetOneBotMessageByDes)

	route.Post("/", Services.CreateOneBotMessage)

	route.Put("/id/:id", Services.UpdateOneBotById)

	route.Delete("/id/:id", Services.DeleteOneBotById)

}
