package main

import (
	"mf-aoc-service/DB"
	"mf-aoc-service/Routes"

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
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Channel-Services"})
	})

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	// }))

	api := app.Group("/api")
	Routes.ChannelRoute(api.Group("/channel"))

	api = app.Group("/api")
	Routes.AdminRoute(api.Group("/admin"))

	api = app.Group("/api")
	Routes.OrgRoute(api.Group("/organization"))

	app.Listen(":3010")

}
