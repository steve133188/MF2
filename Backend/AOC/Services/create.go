package Services

import (
	"fmt"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
)

func AddChannel(c *fiber.Ctx) error {
	usersCollection := DB.MI.ChanDBCol

	data := new(Model.Channel)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate ID in POST")
	}
	data.ID = id.String()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	result, err := usersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.Channel{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func AddAdmin(c *fiber.Ctx) error {
	collection := DB.MI.AdminDBCol

	data := new(Model.Admin)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	data.CreatedTime = time.Now()
	data.UpdatedTime = time.Now()
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate ID for admin")
	}
	data.ID = id.String()
	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert admin",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Admin{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"admin": todo,
		},
	})
}

func AddAgent(c *fiber.Ctx) error {
	usersCollection := DB.MI.OrgDBCol

	data := new(Model.Organization)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate ID in POST")
	}
	data.ID = id.String()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	result, err := usersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.Organization{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
