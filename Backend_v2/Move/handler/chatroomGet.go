package handler

import (
	"Move/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

func GetChatrooms(c *fiber.Ctx) error {

	ctx := context.Background()
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result := make([]*redis.StringStringMapCmd, 0)

	start := time.Now()

	keys, err := RedisClient.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Println(err)
	}

	chatroom := make([]model.ChatroomRedis, 0)

	pipe := RedisClient.Pipeline()
	for _, v := range keys {
		result = append(result, pipe.HGetAll(ctx, v))
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range result {
		data, err := v.Result()
		if err != nil {
			fmt.Println(err)
		}
		temp := new(model.ChatroomRedis)
		err = mapstructure.Decode(data, &temp)
		temp.BotOn, err = strconv.ParseBool(data["BotOn"])
		if err != nil {
			fmt.Println(err)
		}
		temp.IsPin, err = strconv.ParseBool(data["IsPin"])
		if err != nil {
			fmt.Println(err)
		}
		chatroom = append(chatroom, *temp)

	}

	end := time.Since(start)
	fmt.Println("Time: ", end)

	return c.Status(fiber.StatusOK).JSON(chatroom)

}

func GetChatroomsByUser(c *fiber.Ctx) error {

	userID := c.Params("userId")
	filterKey := "Chatroom:*-" + userID

	ctx := context.Background()
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result := make([]*redis.StringStringMapCmd, 0)

	start := time.Now()

	keys, err := RedisClient.Keys(ctx, filterKey).Result()
	if err != nil {
		fmt.Println(err)
	}

	chatroom := make([]model.ChatroomRedis, 0)

	pipe := RedisClient.Pipeline()
	for _, v := range keys {
		result = append(result, pipe.HGetAll(ctx, v))
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range result {
		data, err := v.Result()
		if err != nil {
			fmt.Println(err)
		}
		temp := new(model.ChatroomRedis)
		err = mapstructure.Decode(data, &temp)
		temp.BotOn, err = strconv.ParseBool(data["BotOn"])
		if err != nil {
			fmt.Println(err)
		}
		temp.IsPin, err = strconv.ParseBool(data["IsPin"])
		if err != nil {
			fmt.Println(err)
		}
		chatroom = append(chatroom, *temp)

	}

	end := time.Since(start)
	fmt.Println("Time: ", end)

	return c.Status(fiber.StatusOK).JSON(chatroom)

}
