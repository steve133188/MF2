package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddRoleToUser(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	data := new(Model.Roles)
	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddRoleToUser parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	phone := data.Phone

	filter := bson.D{{"phone", phone}}
	user := new(Model.User)
	err = col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("AddRoleToUser FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if user.Role != "" {
		log.Println("AddRoleToUser User existed role", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user.Role = data.Role
	user.Authority = data.Auth

	res, err := col.UpdateOne(c.Context(), filter, bson.D{{"$set", &user}})
	if err != nil {
		log.Println("AddRoleToUser UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("AddRoleToUser Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUserRoleByPhone(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)
	phone := c.Params("phone")

	auth := Model.Auth{
		Dashboard:        false,
		Livechat:         false,
		Contact:          false,
		Broadcast:        false,
		Flowbuilder:      false,
		Integrations:     false,
		ProductCatalogue: false,
		Organization:     false,
		Admin:            false,
	}

	filter := bson.D{{"phone", phone}}

	err := col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("DeleteUserRole FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.Role = ""
	user.Authority = auth

	res, err := col.UpdateOne(c.Context(), filter, bson.D{{"$set", &user}})
	if res.ModifiedCount == 0 {
		log.Println("DeleteUserRole Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUsersRoleByRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	role := c.Params("role")

	auth := Model.Auth{
		Dashboard:        false,
		Livechat:         false,
		Contact:          false,
		Broadcast:        false,
		Flowbuilder:      false,
		Integrations:     false,
		ProductCatalogue: false,
		Organization:     false,
		Admin:            false,
	}

	filter := bson.D{{"role", role}}

	update := bson.D{
		{"$set", bson.D{
			{"role", ""},
			{"authority", auth},
		}},
	}

	res, err := col.UpdateMany(c.Context(), filter, update)
	if res.ModifiedCount == 0 {
		log.Println("DeleteUserRole Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res.ModifiedCount)
}

func UpdateUserRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)

	data := new(Model.Roles)
	err := c.BodyParser(&data)
	if err != nil {
		if err != nil {
			log.Println("UpdateUserRole parse ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	phone := data.Phone
	role := data.Role
	auth := data.Auth

	filter := bson.D{{"phone", phone}}

	update := bson.D{
		{"$set", bson.D{
			{"role", role},
			{"authority", auth},
		}},
	}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("UpdateUserRole UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("UpdateUserRole Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("UpdateUserRole FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUsersRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var users []Model.User = make([]Model.User, 0)

	data := new(Model.Roles)
	err := c.BodyParser(&data)
	if err != nil {
		if err != nil {
			log.Println("UpdateUsersRole parse ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	role := data.Role
	auth := data.Auth

	filter := bson.D{{"role", role}}

	update := bson.D{
		{"$set", bson.D{
			{"role", role},
			{"authority", auth},
		}},
	}

	res, err := col.UpdateMany(c.Context(), filter, update)
	if err != nil {
		log.Println("UpdateUsersRole UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("UpdateUsersRole Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("UpdateUsersRole Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &users)
	if err != nil {
		log.Println("UpdateUsersRole All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUsersByRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	role := c.Params("role")

	var users []Model.User = make([]Model.User, 0)

	filter := bson.D{{"role", role}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetUsersByRole Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &users)
	if err != nil {
		log.Println("GetUsersByRole All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUserAuthByPhone(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)
	phone := c.Params("phone")

	filter := bson.D{{"phone", phone}}

	err := col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("GetUserAuthByPhone FindOne decode", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user.Authority)
}

func GetUserNumByRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	role := c.Params("role")

	filter := bson.D{{"role", role}}

	count, err := col.CountDocuments(c.Context(), filter)
	if err != nil {
		log.Println("GetUserNumByRole Count ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(count)
}
