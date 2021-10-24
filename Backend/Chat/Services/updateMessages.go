package Services

import (
	"fmt"
	"log"
	"mf-chat-services/DB"
	"mf-chat-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOneMessageById(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Message)
	if err := c.BodyParser(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse message body",
			"error":   err,
		})
	}
	fmt.Println(todo)

	update := bson.D{{Key: "$set", Value: todo}}
	fmt.Println(update)
	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "old_id", Value: c.Params("id")}}, update)
	fmt.Println(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Message failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"message": todo,
		},
	})
}
