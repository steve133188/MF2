package Services

import (
	"fmt"
	"mf-admin-services/DB"
	"mf-admin-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllAdmins(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	// Query to filter
	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.Admin = make([]Model.Admin, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}
	// for i := range todos {
	// 	todos[i].Date = todos[i].Date.Add(time.Hour * 8)
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"records": todos,
		},
	})
}

func GetAdminById(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	// get parameter value
	// paramID, err := primitive.ObjectIDFromHex(c.Params("id")) //valid id: 24 hex
	// if err != nil {
	// 	fmt.Println(err)
	// }

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.Admin{}

	query := bson.D{{Key: "_id", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Analysis record Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"admin": todo,
		},
	})
}
