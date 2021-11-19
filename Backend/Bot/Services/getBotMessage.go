package Services

import (
	"fmt"
	"mf-bot-services/DB"
	"mf-bot-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBots(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	// Query to filter
	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find message",
			"error":   err.Error(),
		})
	}

	var todos []Model.BotBody = make([]Model.BotBody, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into messages",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"messages": todos,
		},
	})
}

func GetOneBotMessageByDes(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	paramID := c.Params("descriotion")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.BotBody{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// timezone UTC +8
	// chat.Time = chat.Time.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Bot Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"bot": todo,
		},
	})
}
