package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf-message-services/Routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Channel-Services"})
	})

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	// }))

	api := app.Group("/api")
	Routes.MessageRoutes(api.Group("/message"))

	app.Listen(":3010")

}
