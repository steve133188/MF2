package Services

import (
	"mf-bot-services/DB"
	"mf-bot-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOneBotMessage(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.BotBody)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	data.CreatedOn = time.Now().Format("January 2, 2006")
	data.UpdatedOn = time.Now().Format("January 2, 2006")
	id := xid.New()
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
	todo := &Model.BotBody{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"bot": todo,
		},
	})
}
