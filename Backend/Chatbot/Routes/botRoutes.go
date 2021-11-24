package Routes

import (
	"mf-bot-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BotBuildRoute(route fiber.Router) {
	//route.Get("/", Services.GetAllBots)

	route.Get("/message/id/:id", Services.GetOneBotMessageByID)
	route.Post("/message/", Services.CreateOneBotMessage)
	route.Put("/message/", Services.UpdateOneBotMessageById)
	route.Delete("/message/id/:id", Services.DeleteOneBotMessageById)

}
