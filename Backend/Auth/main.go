package main

import (
	"mf-auth-servies/DB"
	"mf-auth-servies/Services"
	"mf-auth-servies/Util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	DB.MongoConnect()

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))

	app.Get("/authcheck", Services.AuthCheck)
	app.Post("/login", Services.Login)

	app.Listen(":3008")
}
