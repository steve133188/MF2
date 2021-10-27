package Services

import (
	"mf-bot-services/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOneBotById(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "_id", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Bot Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete bot",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Delete bot (ID = " + paramID + ") success",
	})
}
