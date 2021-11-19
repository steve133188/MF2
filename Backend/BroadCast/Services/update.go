package Services

import (
	"log"
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBroadCastByID(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	todo := new(Model.BroadCast)

	if err := c.BodyParser(todo); err != nil {
		log.Println("UpdateBroadCastByID parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{"name", todo.Name}}, update)
	if err != nil {
		log.Println("UpdateBroadCastByID UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}
