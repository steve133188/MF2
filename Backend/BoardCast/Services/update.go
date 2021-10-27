package Services

import (
	"fmt"
	"log"
	"mf-boardCast-services/DB"
	"mf-boardCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBoardCastByID(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.BoardCast)

	if err := c.BodyParser(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// todo.UpdatedTime = time.Now()
	todo.Id = c.Params("id")
	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "_id", Value: c.Params("id")}}, update)
	fmt.Println(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Board cast failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"boardCast": todo,
		},
	})
}
