package main

import (
	"mf-log-servies/DB"
	"mf-log-servies/Routes"
	"mf-log-servies/Util"

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
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Logs-Services"})
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))
	api := app.Group("/api")

	Routes.CustomersRoute(api.Group("/logs"))

	app.Listen(":3002")

}
