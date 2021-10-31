package Services

import (
	"fmt"
	"mf-customer-services/DB"
	"mf-customer-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteCustomerById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "id", Value: paramID}}

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

func DeleteCustomer(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get param
	data := new(Model.Sort)
	_ = c.BodyParser(&data)
	fmt.Println(data.Name)

	// find todo and return
	filter := bson.D{{Key: "name", Value: data.Name}}

	// find and delete todo

	res, err := customerCollection.DeleteMany(c.Context(), filter)

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
		"success": res,
	})
}
