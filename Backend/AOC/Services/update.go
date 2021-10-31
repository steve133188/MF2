package Services

import (
	"fmt"
	"log"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateChannelById(c *fiber.Ctx) error {
	customersCollection := DB.MI.ChanDBCol
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

func UpdateAdminByID(c *fiber.Ctx) error {
	collection := DB.MI.AdminDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Admin)

	if err := c.BodyParser(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	todo.UpdatedTime = time.Now().Format("January 2, 2006")
	todo.ID = c.Params("id")
	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: c.Params("id")}}, update)
	fmt.Println(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Admin failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"admin": todo,
		},
	})
}

func UpdateOraganizationByPhone(c *fiber.Ctx) error {
	usersCollection := DB.MI.OrgDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.Organization)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	user.ID = c.Params("phone")
	update := bson.D{{Key: "$set", Value: user}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "phone", Value: c.Params("phone")}}, update)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
