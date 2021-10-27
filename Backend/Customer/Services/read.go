package Services

import (
	"fmt"
	"mf-customer-services/DB"
	"mf-customer-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCustomers(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol

	// Query to filter
	query := bson.D{{}}

	cursor, err := customersCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var customers []Model.Customer = make([]Model.Customer, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	// for i := range customers {
	// 	timezone, err := strconv.ParseInt(customers[i].TimeZone, 10, 64)
	// 	if err != nil {
	// 		fmt.Println("Failed to parse TimeZone to int64")
	// 	}
	// 	customers[i].AccountCreatedTime = customers[i].AccountCreatedTime.Add(time.Hour * time.Duration(timezone))
	// 	customers[i].LastUpdatedTime = customers[i].LastUpdatedTime.Add(time.Hour * time.Duration(timezone))
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customers": customers,
		},
	})
}

func GetCustomersById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	// paramID, err := primitive.ObjectIDFromHex(c.Params("id")) //valid id: 24 hex
	// if err != nil {
	// 	fmt.Println(err)
	// }

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.Customer{}

	query := bson.D{{Key: "_id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	// timezone, err := strconv.ParseInt(customer.TimeZone, 10, 64)
	// if err != nil {
	// 	fmt.Println("Failed to parse TimeZone to int64, using default Time Zone UTC +8")
	// 	timezone = 8
	// }

	// customer.AccountCreatedTime = customer.AccountCreatedTime.Add(time.Hour * time.Duration(timezone))
	// customer.LastUpdatedTime = customer.LastUpdatedTime.Add(time.Hour * time.Duration(timezone))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}
