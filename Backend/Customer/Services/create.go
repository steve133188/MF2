package Services

import (
	"fmt"
	"mf-customer-services/DB"
	"mf-customer-services/Model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
)

func AddCustomer(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol

	data := new(Model.Customer)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate UUID")
	}

	data.Id = id.String()
	data.LastUpdatedTime = time.Now()
	data.AccountCreatedTime = time.Now()
	if data.TimeZone == "" {
		data.TimeZone = strconv.FormatInt(8, 10)
	}
	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer",
			"error":   err,
		})
	}

	// get the inserted data
	customer := &Model.Customer{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}
