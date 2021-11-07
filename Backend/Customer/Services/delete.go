package Services

import (
	"mf-customer-services/DB"
	"mf-customer-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteCustomerById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	for _, v := range data.Data {
		query := bson.D{{"id", v}}

		err = customerCollection.FindOneAndDelete(c.Context(), query).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"success": false,
					"message": "Customer Not found",
					"error":   err.Error(),
				})
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot delete customer",
				"error":   err,
			})
		}

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
