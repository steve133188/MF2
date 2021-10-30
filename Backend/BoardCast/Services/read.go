package Services

import (
	"fmt"
	"mf-boardCast-services/DB"
	"mf-boardCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBoardCasts(c *fiber.Ctx) error {
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

	var todos []Model.BoardCast = make([]Model.BoardCast, 0)

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
		"data":    todos,
	})
}

func GetBoardCastsByGroup(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	// Query to filter
	paramID := c.Params("group")
	query := bson.D{{Key: "group", Value: paramID}}

	cursor, err := collection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.BoardCast = make([]Model.BoardCast, 0)

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
		"data":    todos,
	})
}

func GetBoardCastById(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.BoardCast{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Chat Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func GetBoardCastByName(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	paramID := c.Params("name")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.BoardCast{}

	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Chat Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}
