package Services

import (
	"fmt"
	"math/rand"
	"mf-customer-services/DB"
	"mf-customer-services/Model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddCustomer(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol

	data := new(Model.Customer)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	// id := xid.New()

	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int())

	data.ID = id
	data.UpdatedAt = time.Now()
	data.CreatedAt = time.Now()

	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer",
			"error":   err.Error(),
		})
	}

	// get the inserted data
	customer := &Model.Customer{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(customer)
}

func AddManyCustomer(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	// var datas []Model.User = make([]Model.User, 0)
	type data []interface{}
	var datas data
	err := c.BodyParser(&datas)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	// err := json.Unmarshal(c.Body(), &datas)
	_, err = usersCollection.InsertMany(c.Context(), datas)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert agent",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

func AddTags(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol

	data := new(Model.Customer)
	exist := new(Model.Customer)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	err = customersCollection.FindOne(c.Context(), bson.D{{"id", data.ID}}).Decode(exist)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}
	// var val bson.A

	if exist.Tags == nil {
		fmt.Println(1)
		update := bson.D{{"$set", bson.D{{"tags", data.Tags}}}}
		_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: data.ID}}, update)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to update",
				"error":   err.Error(),
			})
		}

	} else {
		fmt.Println(2)
		update := bson.D{{"$addToSet", bson.D{{"tags", bson.D{{"$each", data.Tags}}}}}}

		_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: data.ID}}, update)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to update",
				"error":   err.Error(),
			})
		}

	}

	err = customersCollection.FindOne(c.Context(), bson.D{{"id", data.ID}}).Decode(&exist)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(exist)
}
