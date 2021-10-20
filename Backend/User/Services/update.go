// package Services

// type UpdateUser interface {
// 	UpdateOneUser(obj interface{}) (status string,err error)
// 	UpdateManyUser(obj interface{}) (status string,err error)
// }

// type Update struct {

// }

// func NewUpdate() *Update{
// 	return &Update{}
// }

// func UpdateOneUser(obj interface{}) (status string,err error) {
// 	return
// }

// func UpdateManyUser(obj interface{}) (status string,err error) {
// 	return
// }
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

	update := bson.D{{Key: "$set", Value: user}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: c.Params("id")}}, update)
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
