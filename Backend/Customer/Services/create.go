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
			"error":   err.Error(),
		})
	}

	id := xid.New()

	data.ID = id.String()
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

	isExisted := customersCollection.FindOne(c.Context(), bson.D{{Key: "phone", Value: data.Phone}}).Decode(exist)
	// var val bson.A
	if (isExisted) == nil { //existing division

		update := bson.D{{"$addToSet", bson.D{{"tags", bson.D{{"$each", data.Tags}}}}}}
		_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "phone", Value: data.Phone}}, update)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to update",
				"error":   err.Error(),
			})
		}

		err = customersCollection.FindOne(c.Context(), bson.D{{"phone", data.Phone}}).Decode(&exist)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Customer Not found",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"data":    exist,
		})

	}

	data.ID = xid.New().String()
	_, err = customersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err.Error(),
		})
	}

	err = customersCollection.FindOne(c.Context(), bson.D{{"phone", data.Phone}}).Decode(&exist)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    exist,
	})
}
