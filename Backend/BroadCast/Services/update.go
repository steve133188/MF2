package Services

import (
	"log"
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBroadCastByID(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	todo := new(Model.BroadCast)

	if err := c.BodyParser(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{"name", todo.Name}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Board cast failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}
