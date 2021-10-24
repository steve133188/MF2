package Services

import (
	"mf-chat-services/DB"
	"mf-chat-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddOneMessage(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Message)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// id, err := uuid.NewV4()
	// if err != nil {
	// 	fmt.Println("Failed to generate UUID for message")
	// }
	// data.Id = primitive.NewObjectID()

	// data.DateTime = time.Now()

	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot add message",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Message{}
	query := bson.D{{Key: "old_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"message": todo,
		},
	})
}
