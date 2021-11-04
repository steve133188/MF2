package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func AddRole(c *fiber.Ctx) error {
	collection := DB.MI.AdminDBCol

	data := new(Model.Role)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()
	data.ID = id.String()
	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Role{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func CreateTeam(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Division)
	exist := new(Model.Division)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	isExisted := customersCollection.FindOne(c.Context(), bson.D{{Key: "name", Value: data.Name}}).Decode(exist)
	iVal, _ := json.Marshal(data.Team)
	fmt.Println(bytes.NewBuffer(iVal))
	eVal, _ := json.Marshal(exist.Team)
	fmt.Println(bytes.NewBuffer(eVal))
	if (isExisted) == nil { //existing division
		for _, iv := range iVal {
			for _, ev := range eVal {
				if ev == iv {
					break
				}
				eVal = append(eVal, iv)
			}

		}
		fmt.Println(bytes.NewBuffer(eVal))
	}

	// update := bson.M{"$set": bson.M{"team": data.Team}}

	// isExisted = customersCollection.FindOneAndUpdate(c.Context(), bson.D{{Key: "division", Value: data.Name}}, update).Decode(exist)
	// if (isExisted) == nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "second Is exist",
	// 	})
	// }

	id := xid.New()

	data.TeamId = id.String()
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

func CreateDivision(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Division)
	exist := new(Model.Division)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	isExisted := customersCollection.FindOne(c.Context(), bson.D{{Key: "name", Value: data.Name}}).Decode(exist)
	fmt.Println(isExisted)
	if (isExisted) == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Is exist",
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
