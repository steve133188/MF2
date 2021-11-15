package Services

import (
	"log"
	"mf-customer-services/DB"
	"mf-customer-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCustomerByGroup(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	group := c.Params("group")

	var customers []Model.Customer = make([]Model.Customer, 0)

	query := bson.D{{Key: "group", Value: group}}

	cursor, err := customerCollection.Find(c.Context(), query)
	if err != nil {
		log.Println("GetAllCustomerByGroup find ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		log.Println("GetAllCustomerByGroup all ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	defer cursor.Close(c.Context())
	return c.Status(fiber.StatusOK).JSON(customers)
}

func AddGroupToCustomer(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data struct {
		Group string   `json:"group"`
		Name  []string `json:"name"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddGroupToCustomer all ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	group := data.Group

	query := bson.D{
		{"$and", bson.A{
			bson.D{{"name", bson.D{{"$in", data.Name}}}},
			bson.D{{"group", ""}},
		}},
	}
	update := bson.D{{"$set", bson.D{{"group", group}}}}

	_, err = col.UpdateMany(c.Context(), query, update)
	if err != nil {
		log.Println("AddGroupToCustomer UpdateMany ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	var result []Model.Customer = make([]Model.Customer, 0)
	cursor, err := col.Find(c.Context(), bson.D{{"name", bson.D{{"$in", data.Name}}}})
	if err != nil {
		log.Println("AddGroupToCustomer Find ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &result)
	if err != nil {
		log.Println("AddGroupToCustomer All ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func EditGroupName(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data struct {
		Old    string `json:"old"`
		Update string `json:"update"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("EditGroupName parse ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"group", data.Old}}
	update := bson.D{{"$set", bson.D{{"group", data.Update}}}}

	updateResult, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("EditGroupName UpdateMany ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}
	if updateResult.ModifiedCount > 0 {
		var results []Model.Customer = make([]Model.Customer, 0)

		cursor, err := col.Find(c.Context(), bson.D{{"group", data.Update}})
		if err != nil {
			log.Println("EditGroupName Find ", err)
			c.SendStatus(fiber.StatusInternalServerError)
		}

		err = cursor.All(c.Context(), &results)
		if err != nil {
			log.Println("EditGroupName All ", err)
			c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(results)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}

//404
func DeleteGrpupByName(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	group := c.Params("group")

	filter := bson.D{{"group", group}}
	update := bson.D{{"$set", bson.D{{"group", ""}}}}

	updateResult, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("DeleteGrpupByName AUpdateMany ", err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	if updateResult.ModifiedCount > 0 {
		var results []Model.Customer = make([]Model.Customer, 0)

		cursor, err := col.Find(c.Context(), bson.D{{"group", group}})
		if err != nil {
			log.Println("DeleteGrpupByName Find ", err)
			c.SendStatus(fiber.StatusInternalServerError)
		}

		err = cursor.All(c.Context(), &results)
		if err != nil {
			log.Println("DeleteGrpupByName All ", err)
			c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(results)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}

//GetGroupOfCustomers

//GetGroups
