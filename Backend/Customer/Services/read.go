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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

func GetCustomersById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := new(Model.Customer)

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

func GetCustomersByName(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	paramID := c.Params("name")
	fmt.Println(paramID)

	// find todo and return
	customer := new(Model.Customer)

	query := bson.D{{Key: "name", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

func GetAllByTeamSorting(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	paramID := c.Params("team")
	fmt.Println(paramID)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)

	query := bson.D{{Key: "team", Value: paramID}}

	cursor, err := customerCollection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

func GetAgentSorting(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)
	// fmt.Println(data.Name)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Name {
		val = append(val, v)
	}

	filter := bson.D{{"agent", bson.D{{"$in", val}}}}
	fmt.Println(filter)
	cursor, err := customerCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

func GetTagsSorting(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Name {
		val = append(val, v)
	}

	filter := bson.D{{"tags", bson.D{{"$in", val}}}}
	fmt.Println(filter)
	// filter := bson.D{{Key: "tags", Value: data.Name}}
	cursor, err := customerCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

func GetChannelSorting(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)
	fmt.Println(data.Name)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Name {
		val = append(val, v)
	}

	filter := bson.D{{"channel", bson.D{{"$in", val}}}}
	fmt.Println(filter)
	// filter := bson.D{{Key: "channel", Value: data.Name}}

	cursor, err := customerCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}
