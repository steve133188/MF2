package main

import (
	"ChatbotAPI/config"
	"ChatbotAPI/hanlder"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
	"reflect"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.RedisInit()

	go func() {

		messagesSub := config.ChatBotDB.Subscribe(context.Background(),"messages.received" )

		for{
			for msg := range messagesSub.Channel() {

				data ,err :=json.Marshal(msg.Payload)

				fmt.Println(msg.Pattern)
				fmt.Println(msg.PayloadSlice)
				fmt.Println(msg.String())
				fmt.Println(msg.Payload)

				dataMap :=map[string]interface{}{

				}
				if err != nil{
					fmt.Printf("Marshall failed=%s\n",err)
				}
				err = json.Unmarshal(data ,&dataMap);if err != nil {
					fmt.Printf("UNMarshall failed=%s\n",err)
				}

				for i ,v :=range dataMap{
					fmt.Println( i, v)

				}

				fmt.Println( dataMap, reflect.TypeOf(dataMap))
				//config.ChatBotDB.Set(context.Background() , key , )

			}
		}

	}()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Chat Bot API Server is running")
	})

	api := app.Group("/api")

	api.Post("/flow", hanlder.CreateFlow)
	api.Post("/action", hanlder.CreateAction)
	api.Post("/option", hanlder.CreateOption)

	app.Post("/" , func(c *fiber.Ctx) error {

		config.ChatBotDB.Publish(context.Background() ,"messages.received" , c.Body())

		return c.Status(http.StatusOK).SendString("received data")
	})

	app.Listen(":3010")

}
