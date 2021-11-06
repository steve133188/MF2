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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	customer.UpdatedAt = time.Now()

	update := bson.D{{Key: "$set", Value: &customer}}

	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: customer.ID}}, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Customer failed to update",
			"error":   err.Error(),
		})
	}

	query := bson.D{{Key: "id", Value: customer.ID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// func UpdateCustomerTags(c *fiber.Ctx) error {
// 	customersCollection := DB.MI.DBCol
// 	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	customer := new(Model.Customer)

// 	if err := c.BodyParser(&customer); err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Failed to parse body",
// 			"error":   err,
// 		})
// 	}

// 	customer.UpdatedAt = time.Now()

// 	update := bson.M{"$set": bson.M{
// 		"tags": customer.Tags,
// 	}}
// 	_, err := customersCollection.FindOneAndUpdate(c.Context(), bson.D{{"phone", c.Params("name")}}, update)
// 	fmt.Println(customer)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Customer failed to update",
// 			"error":   err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 		"data":    customer,
// 	})
// }

func UpdateChannelInfoByPhone(c *fiber.Ctx) error {
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

	customer.UpdatedAt = time.Now()

	update := bson.M{"$set": bson.M{
		"channel_info": customer.ChannelInfo,
	}}
	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "phone", Value: c.Params("phone")}}, update)
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
