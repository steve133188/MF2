package Services

import (
	"fmt"
	"mf-flowbuilder-services/DB"
	"mf-flowbuilder-services/Model"

	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOneBot(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Bot)

	err := c.BodyParser(&data)
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
		fmt.Println("Failed to generate bot id")
	}
	data.ID = id.String()

	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert bot",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Bot{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"bot": todo,
		},
	})
}
