package main

import (
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
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Users-Services"})
	})

	app.Post("/api/users/login", Services.Login)
	app.Post("/api/users", Services.AddAgent)
	app.Post("/api/users/addMany", Services.AddManyAgent)
	app.Post("/api/users/forgot-password", Services.ForgotPassword)
	// app.Post("/api/users/forgot-password", Services.ForgotPassword)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	}))
	api := app.Group("/api")

	Routes.UsersRoute(api.Group("/users"))
	Routes.RoleRoute(api.Group("/roles"))

	app.Listen(":3001")
}
