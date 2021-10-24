package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	data := new(Model.User)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// id, err := uuid.NewV4()
	// if err != nil {
	// 	fmt.Println("Failed to generate ID in POST")
	// }
	// data.ID = id.String()
	// data.Date = time.Now()
	result, err := usersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.User{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	fmt.Println(result.InsertedID)
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}
