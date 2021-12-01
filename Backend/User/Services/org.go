package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTeamIDToUser(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var data struct {
		UserPhone string `json:"user_phone" bson:"user_phone"`
		TeamID    string `json:"team_id" bson:"team_id"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddTeanIDToUser parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"phone", data.UserPhone}}
	update := bson.D{{"$set", bson.D{{"team_id", data.TeamID}}}}

	result := new(Model.User)
	err = col.FindOne(c.Context(), bson.D{{"phone", data.UserPhone}}).Decode(&result)
	if err != nil {
		log.Println("AddTeanIDToUser FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result.TeamID != "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err = col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("AddTeanIDToUser UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{"phone", data.UserPhone}}).Decode(&result)
	if err != nil {
		log.Println("AddTeanIDToUser FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateUserTeam(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var data struct {
		UserPhone string `json:"user_phone" bson:"user_phone"`
		TeamID    string `json:"team_id" bson:"team_id"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("UpdateUserTeam parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"phone", data.UserPhone}}
	update := bson.D{{"$set", bson.D{{"team_id", data.TeamID}}}}

	result := new(Model.User)
	_, err = col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("AddTeanIDToUser UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{"phone", data.UserPhone}}).Decode(&result)
	if err != nil {
		log.Println("AddTeanIDToUser FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateUsersTeamID(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var data struct {
		OldId string `json:"old_id" bson:"old_id"`
		NewId string `json:"new_id" bson:"new_id"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddTeanIDToUser parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"team_id", data.OldId}}
	update := bson.D{{"$set", bson.D{{"team_id", data.NewId}}}}

	_, err = col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("UpdateUsersTeamID UpdateMany ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var results []Model.User = make([]Model.User, 0)
	cursor, err := col.Find(c.Context(), bson.D{{"team_id", data.NewId}})
	if err != nil {
		log.Println("UpdateUsersTeamID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &results)
	if err != nil {
		log.Println("UpdateUsersTeamID All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func DeleteTeamIDFromUsers(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	teamId := c.Params("team")

	filter := bson.D{{"team_id", teamId}}
	update := bson.D{{"$set", bson.D{{"team_id", ""}}}}

	_, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("DeleteTeamIDfromUser UpdateMany ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUsersByTeamID(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	teamId := c.Params("id")

	filter := bson.D{{"team_id", teamId}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetUsersByTeamID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var results []Model.User = make([]Model.User, 0)

	err = cursor.All(c.Context(), &results)
	if err != nil {
		log.Println("GetUsersByTeamID All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func GetUserWithNoTeam(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	filter := bson.D{{"team_id", ""}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetUserWithNoTeam Find     ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var users []Model.User = make([]Model.User, 0)

	err = cursor.All(c.Context(), &users)
	if err != nil {
		log.Println("GetUserWithNoTeam All     ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
