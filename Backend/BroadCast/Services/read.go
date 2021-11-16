package Services

import (
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBroadCasts(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetBroadCastsByGroup(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Param)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	query := bson.D{{"group", data.Param}}

	cursor, err := collection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetBroadCastsByName(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Param)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	query := bson.D{{"name", data.Param}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}
