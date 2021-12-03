package Services

import (
	"log"
	"mf-customer-services/DB"
	"mf-customer-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCustomerByID(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	customer := new(Model.Customer)

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	customer.UpdatedAt = time.Now()

	update := bson.D{{Key: "$set", Value: &customer}}

	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: customer.ID}}, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Customer failed to update",
			"error":   err.Error(),
		})
	}

	query := bson.D{{Key: "id", Value: customer.ID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(customer)
}

func UpdateCustomersTags(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var data struct {
		New string `json:"new" bson:"new"`
		Old string `json:"old" bson:"old"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	query := bson.D{{"tags", data.Old}}
	update := bson.D{{"$set", bson.D{{"tags.$", data.New}}}}

	_, err := customersCollection.UpdateMany(c.Context(), query, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}

	cursor, err := customersCollection.Find(c.Context(), bson.D{{"tags", data.New}})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	var customers []Model.Customer = make([]Model.Customer, 0)

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(customers)
}

func DeleteTagFromAllCustomer(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var data struct {
		Old string `json:"old" bson:"old"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	query := bson.D{{"tags", data.Old}}
	update := bson.D{{"$pull", bson.D{{"tags", data.Old}}}}

	result, err := customersCollection.UpdateMany(c.Context(), query, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func DeleteCustomerTags(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := new(Model.Customer)
	var data struct {
		ID  string   `json:"id" bson:"id"`
		Old []string `json:"old" bson:"old"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	query := bson.D{{"id", data.ID}}
	update := bson.D{{"$pull", bson.D{{"tags", bson.D{{"$in", data.Old}}}}}}

	_, err := customersCollection.UpdateOne(c.Context(), query, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	err = customersCollection.FindOne(c.Context(), bson.D{{"id", data.ID}}).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Find",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func UpdateManyCustomers(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var customers []Model.Customer = make([]Model.Customer, 0)

	var results []Model.Customer = make([]Model.Customer, 0)
	result := new(Model.Customer)

	err := c.BodyParser(&customers)
	if err != nil {
		log.Println("UpdateManyCustomers parse     ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	for _, customer := range customers {
		filter := bson.D{{"id", customer.ID}}
		update := bson.D{{"$set", &customer}}
		res, err := col.UpdateOne(c.Context(), filter, update)
		if err != nil {
			log.Println("UpdateManyCustomers UpdateOne     ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		log.Println(customer.ID, "    ", res.ModifiedCount)

		err = col.FindOne(c.Context(), filter).Decode(&result)
		if err != nil {
			log.Println("UpdateManyCustomers FindOne     ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		results = append(results, *result)

	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func PutPhoneToCustomer(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data struct {
		ID    string `json:"id"`
		Phone string `json:"phone"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("PutPhoneToCustomer parse    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"id", data.ID}}
	update := bson.D{{"$set", bson.D{{"phone", data.Phone}}}}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("PutPhoneToCustomer UpdateOne    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("PutPhoneToCustomer Update Failed    ", err)
		return c.SendStatus(fiber.StatusNotModified)
	}

	customer := new(Model.Customer)

	err = col.FindOne(c.Context(), filter).Decode(&customer)
	if err != nil {
		log.Println("PutPhoneToCustomer FindOne    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
