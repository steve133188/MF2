package Services

import (
	"fmt"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllChannelInfo(c *fiber.Ctx) error {
	fmt.Println("getall")

	usersCollection := DB.MI.ChanDBCol
	// Query to filter
	query := bson.D{{}}

	cursor, err := usersCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var users []Model.Channel = make([]Model.Channel, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func GetChannelInfoById(c *fiber.Ctx) error {
	customerCollection := DB.MI.ChanDBCol

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.Channel{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

func GetAllRole(c *fiber.Ctx) error {
	collection := DB.MI.RoleDBCol

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
	var todos []Model.Role = make([]Model.Role, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todos,
	})
}

func GetAllTags(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

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
	var todos []Model.Tags = make([]Model.Tags, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todos,
	})
}

func GetRoleById(c *fiber.Ctx) error {
	collection := DB.MI.RoleDBCol

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.Role{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func GetRoleByName(c *fiber.Ctx) error {
	collection := DB.MI.RoleDBCol

	paramID := c.Params("name")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.Role{}

	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func GetTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	paramID := c.Params("name")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.Tags{}

	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func GetAllOrgInfo(c *fiber.Ctx) error {
	fmt.Println("getall")
	usersCollection := DB.MI.OrgDBCol
	// Query to filter
	query := bson.D{{}}

	cursor, err := usersCollection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	defer cursor.Close(c.Context())
	var users []Model.Division = make([]Model.Division, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}
