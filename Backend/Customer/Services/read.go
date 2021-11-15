package Services

import (
	"fmt"
	"log"
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
		log.Println("GetAllCustomers Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var customers []Model.Customer = make([]Model.Customer, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &customers)
	if err != nil {
		log.Println("GetAllCustomers All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}

func GetCustomerById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	id := c.Params("id")

	customer := new(Model.Customer)

	query := bson.D{{Key: "id", Value: id}}

	err := customerCollection.FindOne(c.Context(), query).Decode(&customer)
	if err != nil {
		log.Println("GetCustomerById findone ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func GetCustomerByName(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	name := c.Params("name")
	customer := new(Model.Customer)

	query := bson.D{{Key: "name", Value: name}}

	err := customerCollection.FindOne(c.Context(), query).Decode(&customer)
	if err != nil {
		log.Println("GetCustomerByName findone ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func GetChannelInfoByPhone(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	paramID := c.Params("phone")
	fmt.Println(paramID)

	// find todo and return
	customer := new(Model.Customer)

	query := bson.D{{Key: "phone", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func GetAgentFilter(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)
	// fmt.Println(data.Name)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Data {
		val = append(val, v)
	}

	filter := bson.D{{"agents", bson.D{{"$in", val}}}}
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

	return c.Status(fiber.StatusOK).JSON(customers)
}

func GetTagsFilter(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Data {
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

	return c.Status(fiber.StatusOK).JSON(customers)
}

func GetChannelFilter(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Data {
		val = append(val, v)
	}

	filter := bson.D{{"channels", bson.D{{"$in", val}}}}
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

	return c.Status(fiber.StatusOK).JSON(customers)
}

func GetTeamFilter(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	data := new(Model.Sort)
	_ = c.BodyParser(&data)
	// fmt.Println(data.Name)

	// find todo and return
	var customers []Model.Customer = make([]Model.Customer, 0)
	var val bson.A
	for _, v := range data.Data {
		val = append(val, v)
	}

	filter := bson.D{{"team", bson.D{{"$in", val}}}}
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

	return c.Status(fiber.StatusOK).JSON(customers)
}
