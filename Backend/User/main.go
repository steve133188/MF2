package main

import (
	"mf-user-servies/DB"
	"mf-user-servies/Routes"
	"mf-user-servies/Services"

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
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Users-Services"})
	})

	app.Get("/test/add", Services.AddUser)
	app.Get("/all", Services.GetAllUsers)
	app.Post("/login", Services.LoginUser)
	api := app.Group("/api")

	Routes.UsersRoute(api.Group("/users"))

	app.Listen(":3001")
}
