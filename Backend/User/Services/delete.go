// package Services

// type DeleteUser interface {
// 	DeleteOneUser(obj interface{}) (status string,err error)
// 	DeleteManyUser(obj interface{}) (status string,err error)
// 	DeleteAllUser(obj interface{}) (status string,err error)
// }

// type Delete struct {

// }

// func NewDelete() *Delete{
// 	return &Delete{}
// }

// func (d*Delete)DeleteOneUser(obj interface{}) (status string,err error){
// 	 return
// }

// func (d*Delete)DeleteManyUser(obj interface{}) (status string,err error){
// 	 return
// }

// func (d*Delete)DeleteAllUser(obj interface{}) (status string,err error){
// 	 return
// }

package Services

import (
	"mf-user-servies/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUserById(c *fiber.Ctx) error {
	userCollection := DB.MI.DBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "id", Value: paramID}}

	err := userCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete user",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Delete user (ID = " + paramID + ") success",
	})
}
