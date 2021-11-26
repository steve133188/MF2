package Services

import (
	"log"
	"mf-customer-services/DB"
	"mf-customer-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTeamIDToCustomer(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data struct {
		CustomerID string `json:"customer_id" bson:"customer_id"`
		TeamID     string `json:"team_id" bson:"team_id"`
	}

	filter := bson.D{{"id", data.CustomerID}}
	update := bson.D{{"$set", bson.D{{"team_id", data.TeamID}}}}

	result := new(Model.Customer)
	err := col.FindOne(c.Context(), filter).Decode(&result)
	if err != nil {
		log.Println("AddTeamIDToCustomer FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result.TeamID != "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err = col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("AddTeamIDToCustomer UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), filter).Decode(&result)
	if err != nil {
		log.Println("AddTeamIDToCustomer FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateCustomersTeamID(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data struct {
		OldId string `json:"old_id" bson:"old_id"`
		NewId string `json:"new_id" bson:"new_id"`
	}

	filter := bson.D{{"old_id", data.OldId}}
	update := bson.D{{"$set", bson.D{{"team_id", data.NewId}}}}

	_, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("UpdateCustomersTeamID UpdateMany ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var results []Model.Customer = make([]Model.Customer, 0)
	cursor, err := col.Find(c.Context(), bson.D{{"team_id", data.NewId}})
	if err != nil {
		log.Println("UpdateCustomersTeamID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &results)
	if err != nil {
		log.Println("UpdateCustomersTeamID All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func DeleteTeamIDFromCustomers(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	teamId := c.Params("team")

	filter := bson.D{{"team_id", teamId}}
	update := bson.D{{"$set", bson.D{{"team_id", ""}}}}

	_, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("DeleteTeamIDFromUsers UpdateMany ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetCustomersByTeamID(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	teamId := c.Params("id")

	filter := bson.D{{"team_id", teamId}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetCustomersByTeamID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var results []Model.Customer = make([]Model.Customer, 0)

	err = cursor.All(c.Context(), &results)
	if err != nil {
		log.Println("GetCustomersByTeamID All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func GetCustomersWithNoTeam(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	filter := bson.D{{"team_id", ""}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetCustomersWithNoTeam Find     ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var customers []Model.Customer = make([]Model.Customer, 0)

	err = cursor.All(c.Context(), &customers)
	if err != nil {
		log.Println("GetCustomersWithNoTeam All     ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}
