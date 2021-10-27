package Services

import (
	"fmt"
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserByID(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// user.Date = time.Now()
	user.ID = c.Params("id")
	update := bson.D{{Key: "$set", Value: user}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "_id", Value: c.Params("id")}}, update)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}
