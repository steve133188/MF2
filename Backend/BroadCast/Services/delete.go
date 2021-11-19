package Services

import (
	"log"
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteBroadCastByName(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Param)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("DeleteBroadCastByName parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	query := bson.D{{"name", data.Param}}

	err = collection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("DeleteBroadCastByName ErrNoDocuments ", err)
			return c.SendStatus(fiber.StatusNotFound)
		}
		log.Println("DeleteBroadCastByName FindOneAndDelete ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
