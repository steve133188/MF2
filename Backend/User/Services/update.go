package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserByName(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol
	user := new(Model.User)

	if err := c.BodyParser(&user); err != nil {
		log.Println("UpdateUserByName parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	password, err := Util.HashPassword(user.Password)
	if err != nil {
		log.Println("UpdateUserByName hashPW: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.Password = password

	update := bson.D{{Key: "$set", Value: user}}

	_, err = col.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)

	if err != nil {
		log.Println("UpdateUserByName UpdateOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&user)
	if err != nil {
		log.Println("UpdateUserByName FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func ChangeUserStatus(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol
	// user := new(Model.User)
	data := new(Model.User)

	var body struct {
		Username string `json:"username"`
		Status   string `json:"status"`
	}

	err := c.BodyParser(&body)
	if err != nil {
		log.Println("ChangeUserStatus parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	update := bson.M{"$set": bson.M{
		"status": body.Status,
	}}

	_, err = col.UpdateOne(c.Context(), bson.D{{Key: "username", Value: body.Username}}, update)
	if err != nil {
		log.Println("ChangeUserStatus UpdateOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{Key: "username", Value: body.Username}}).Decode(&data)
	if err != nil {
		log.Println("ChangeUserStatus FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

func UpdateChannelInfoByPhone(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol
	chanInfo := new(Model.Info)
	phone := c.Params("phone")
	update := bson.D{{"$set", bson.D{{"channel_info", chanInfo}}}}

	_, err := col.UpdateOne(c.Context(), bson.D{{Key: "phone", Value: phone}}, update)

	if err != nil {
		log.Println("UpdateChannelInfoByPhone UpdateOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	customer := new(Model.User)
	col.FindOne(c.Context(), bson.D{{"phone", chanInfo.Phone}}).Decode(&customer)
	return c.Status(fiber.StatusCreated).JSON(customer)
}
