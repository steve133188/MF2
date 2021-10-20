package Services

import (
	"fmt"
	"mf-admin-services/DB"
	"mf-admin-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
)

func AddAdmin(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Admin)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	data.CreatedTime = time.Now()
	data.UpdatedTime = time.Now()
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate ID for admin")
	}
	data.ID = id.String()
	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert admin",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Admin{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"admin": todo,
		},
	})
}
