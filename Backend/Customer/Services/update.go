package Services

import (
	"fmt"
	"log"
	"mf-customer-services/DB"
	"mf-customer-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCustomerByID(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	customer := new(Model.Customer)

	if err := c.BodyParser(&customer); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// customer.LastUpdatedTime = time.Now()
	customer.ID = c.Params("id")
	customer.UpdatedAt = time.Now().Format("January 2, 2006")
	// if customer.TimeZone == "" {
	// 	customer.TimeZone = strconv.FormatInt(8, 10)
	// }
	update := bson.D{{Key: "$set", Value: &customer}}

	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: c.Params("id")}}, update)
	fmt.Println(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Customer failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}
