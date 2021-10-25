package Services

import (
	"mf-boardCast-services/DB"
	"mf-boardCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddBoardCast(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.BoardCast)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// data.CreatedTime = time.Now()
	// data.UpdatedTime = time.Now()
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	fmt.Println("Failed to generate ID for boardcast")
	// }
	// data.ID = id.String()
	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.BoardCast{}
	query := bson.D{{Key: "old_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"boardCast": todo,
		},
	})
}
