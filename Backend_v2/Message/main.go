package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf2-message-server/config"
	"mf2-message-server/router"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.RedisInit()
	//config.TestRedis()

	//app.Post("/test", func(c *fiber.Ctx) error {
	//	msg := make(map[string]interface{})
	//
	//	input := c.Body()
	//
	//	err := json.Unmarshal(input, &msg)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	message := new(model.Message)
	//
	//	for k, v := range msg {
	//		fmt.Println(k)
	//		val, _ := json.Marshal(v)
	//		json.Unmarshal(val, &message)
	//		fmt.Println(message)
	//	}
	//
	//	return c.Status(fiber.StatusOK).JSON(message)
	//})

	api := app.Group("/api-v2")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Message API Server is running")
	})

	router.MsgRouter(api.Group("/message"))
	router.MsgsRouter(api.Group("/messages"))

	app.Listen(":8080")
}
