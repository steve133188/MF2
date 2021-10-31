package main

import (
	"mf-organization-services/DB"
	"mf-organization-services/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Organization-Services"})
	})

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	// }))

	api := app.Group("/api")

	Routes.CustomersRoute(api.Group("/organization"))

	app.Listen(":3011")

}
