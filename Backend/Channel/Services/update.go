package Services

import (
	"fmt"
	"log"
	"mf-channel-service/DB"
	"mf-channel-service/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateChannelById(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	customer := new(Model.Channel)

	if err := c.BodyParser(customer); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// customer.LastUpdatedTime = time.Now()
	customer.ID = c.Params("id")
	// if customer.TimeZone == "" {
	// 	customer.TimeZone = strconv.FormatInt(8, 10)
	// }
	update := bson.D{{Key: "$set", Value: customer}}

	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: c.Params("id")}}, update)
	fmt.Println(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}
