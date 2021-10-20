package main

import (
	"mf-auth-servies/DB"
	"mf-auth-servies/Services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	DB.MongoConnect()

	app.Get("/authcheck", Services.AuthCheck)
	app.Post("/login", Services.Login)

	app.Listen(":3008")
}
