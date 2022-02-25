package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mf2-channel-router/config"
	"mf2-channel-router/handler"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// redis initialization
	config.RedisInit()
	//config.TestRedis()
	//sub1 := config.TestClient.Subscribe(context.Background(), "messages")
	//go func() {
	//	for msg := range sub1.Channel() {
	//		// 打印收到的消息
	//		fmt.Println(msg.Channel)
	//		fmt.Println(msg.Payload)
	//
	//		resp, err := http.Post("http://localhost:8080/test", "Content-Type: application/json", strings.NewReader(msg.Payload))
	//		if err != nil {
	//			log.Println(err)
	//		}
	//		body, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//		fmt.Println(string(body))
	//	}
	//}()

	api := app.Group("/ch-router")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("MF2 Channel Router is running")
	})

	api.Post("/connect", handler.ChannelConnect)
	api.Get("/restart", handler.ChannelRestart)
	api.Get("/disconnect", handler.ChannelDisconnect)
	api.Post("/send-message", handler.ChannelSendMessage)

	api.Get("/update-status", handler.ChannelUpdateStatus)

	app.Listen(":8081")
}
