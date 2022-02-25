package handler

import (
	"Move/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UpdateChatroomUnreadToZero(c *fiber.Ctx) error {

	ctx := context.Background()

	channel := c.Params("channel")
	roomID := c.Params("room_id")
	filterKey := channel + ":Chatroom:" + roomID

	err := config.RedisClient.HSet(ctx, filterKey, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	chatroom, err := config.RedisClient.HGetAll(ctx, filterKey).Result()
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}
