package Services

import (
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
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

	id := xid.New()
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

	data.CreatedTime = time.Now().Format("January 2, 2006")
	data.UpdatedTime = time.Now().Format("January 2, 2006")
	id := xid.New()
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

func CreateTeam(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Team)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()

	data.ID = id.String()
	data.CreatedAt = time.Now().Format("January 2, 2006")
	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	// get the inserted data
	customer := &Model.Team{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

func CreateDivision(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Division)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()

	data.ID = id.String()
	data.CreatedAt = time.Now().Format("January 2, 2006")
	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	// get the inserted data
	customer := &Model.Division{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}
