package handler

import (
	"Move/config"
	"Move/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func MoveChatroom(c *fiber.Ctx) error {
	chatrooms := make([]model.Chatroom, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pChatrooms := make([]model.Chatroom, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pChatrooms)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		chatrooms = append(chatrooms, pChatrooms...)
	}

	ctx := context.Background()
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Start to fill in Redis")

	for i, v := range chatrooms {
		fmt.Printf("%d ", i)
		temp := structs.Map(&v)
		temp["BotOn"] = strconv.FormatBool(v.BotOn)
		temp["IsPin"] = strconv.FormatBool(v.IsPin)
		err := RedisClient.HMSet(ctx, v.Channel+":Chatroom:"+v.RoomID, temp).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	return c.Status(fiber.StatusOK).JSON(chatrooms)

}
