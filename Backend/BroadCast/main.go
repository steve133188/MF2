package main

import (
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Routes"
	"mf-broadCast-services/Util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-BroadCasts-Services"})
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))

	api := app.Group("/api")

	Routes.BroadCastRoute(api.Group("/broadcasts"))

	app.Listen(":3007")
}
