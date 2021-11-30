package main

import (
	"fmt"
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Routes"
	"mf-user-servies/Services"
	"mf-user-servies/Util"

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
	app.Post("/test", func(c *fiber.Ctx) error {
		var data []interface{}

		err := c.BodyParser(&data)
		if err != nil {
			log.Println("testing parse    ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		fmt.Println(c.JSON(data))

		return c.Status(fiber.StatusOK).JSON(data)
	})

	app.Post("/api/users/login", Services.Login)
	app.Post("/api/users", Services.AddAgent)
	app.Post("/api/users/addMany", Services.AddManyAgent)
	app.Post("/api/users/forgot-password", Services.ForgotPassword)
	app.Put("/api/users/check-auth/:email", Services.CheckUserAuthority)

	// app.Post("/api/users/forgot-password", Services.ForgotPassword)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))
	api := app.Group("/api")

	Routes.UsersRoute(api.Group("/users"))

	app.Listen(":3001")
}
