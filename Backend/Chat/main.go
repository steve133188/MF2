package main

import (
	"mf-chat-services/DB"
	"mf-chat-services/Routes"
	"mf-chat-services/Util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

type test struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Chats-Services"})
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))

	api := app.Group("/api")

	Routes.ChatRoute(api.Group("/messages"))

	app.Listen(":3003")

}
