package Services

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"time"
)

func GetAllUsers(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

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

	var users []Model.User = make([]Model.User, 0)

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
		"data": fiber.Map{
			"users": users,
		},
	})
}

func GetUsersById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.User{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}
func LoginUser(c*fiber.Ctx) error{
	users_collection := DB.MI.DBCol
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user Model.User

	q:= new(Model.User)

	if err := c.BodyParser(&q); err != nil {
		return err
	}

	fmt.Println(q)

	filter := bson.M{"email":q.Email,"password":q.Password}

	findResult :=users_collection.FindOne(ctx , filter)
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	err := findResult.Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})

}