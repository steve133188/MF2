package Services

import (
	"mf-customer-services/DB"
	"mf-customer-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
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
			"error":   err,
		})
	}

	id := xid.New()

	data.ID = id.String()
	data.UpdatedAt = time.Now().Format("January 2, 2006")
	data.CreatedAt = time.Now().Format("January 2, 2006")
	// data.AccountCreatedTime = time.Now()
	// if data.TimeZone == "" {
	// 	data.TimeZone = strconv.FormatInt(8, 10)
	// }
	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer",
			"error":   err,
		})
	}

	// get the inserted data
	customer := &Model.Customer{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
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
