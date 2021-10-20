package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf-auth-servies/Services"
	"mf-auth-servies/DB"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	DB.MongoConnect()

	app.Get("/authcheck", Services.AuthCheck)
	app.Post("/login", Services.Login)

	app.Listen(":3001")
}