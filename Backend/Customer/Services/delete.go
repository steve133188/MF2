package Services

import (
	"mf-customer-services/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteCustomerById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "old_id", Value: paramID}}

	err := customerCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Customer Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete customer",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Delete customer (ID = " + paramID + ") success",
	})
}
