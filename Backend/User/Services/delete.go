package Services

import (
	"log"
	"mf-user-servies/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUserByName(c *fiber.Ctx) error {
	userCollection := DB.MI.UserDBCol

	name := c.Params("name")

	query := bson.M{"username": name}

	err := userCollection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("DeleteUserByName FindOneAndDelete: ", err)
			return c.SendStatus(fiber.StatusNotFound)
		}

		log.Println("DeleteUserByName FindOneAndDelete: ", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
