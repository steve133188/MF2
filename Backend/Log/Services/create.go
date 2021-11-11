package Services

import (
	"mf-log-servies/DB"
	"mf-log-servies/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserLog(c *fiber.Ctx) error {
	logCollection := DB.MI.DBCol

	data := new(Model.UserLog)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()
	data.ID = id.String()
	data.Date = time.Now().Format("January 2 2006 15:04:05")

	result, err := logCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user log",
			"error":   err,
		})
	}

	// get the inserted data
	userLog := &Model.UserLog{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	logCollection.FindOne(c.Context(), query).Decode(userLog)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"userLog": userLog,
		},
	})
}

func CreateCustomerLog(c *fiber.Ctx) error {
	logCollection := DB.MI.DBCol

	data := new(Model.CustomerLog)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	data.Date = time.Now().Format("January 2 2006 15:04:05")

	_, err = logCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer log",
			"error":   err.Error(),
		})
	}

	// get the inserted data
	customerLog := &Model.CustomerLog{}
	query := bson.D{{"customer_name", data.CustomerName}}

	logCollection.FindOne(c.Context(), query).Decode(customerLog)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    customerLog,
	})
}

func CreateSystemLog(c *fiber.Ctx) error {
	logCollection := DB.MI.DBCol

	data := new(Model.SystemLog)

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
	data.Date = time.Now().Format("January 2 2006 15:04:05")

	result, err := logCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert system log",
			"error":   err,
		})
	}

	// get the inserted data
	systemLog := &Model.SystemLog{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	logCollection.FindOne(c.Context(), query).Decode(systemLog)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"systemLog": systemLog,
		},
	})
}

func AddManyUserLog(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	// var datas []Model.User = make([]Model.User, 0)
	type data []interface{}
	var datas data
	err := c.BodyParser(&datas)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	// err := json.Unmarshal(c.Body(), &datas)
	_, err = usersCollection.InsertMany(c.Context(), datas)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert agent",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}
