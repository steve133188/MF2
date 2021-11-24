package main

import (
	"mf-bot-services/DB"
	"mf-bot-services/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	api := app.Group("/api")
	Routes.BotBuildRoute(api.Group("/botMessages"))

	app.Listen(":3005")

}
