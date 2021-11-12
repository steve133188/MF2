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

	name := c.Query("name")
	status := c.Query("status")

	update := bson.M{"$set": bson.M{
		"status": status,
	}}

	_, err := col.UpdateOne(c.Context(), bson.D{{Key: "username", Value: name}}, update)
	if err != nil {
		log.Println("ChangeUserStatus UpdateOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{Key: "username", Value: name}}).Decode(&data)
	if err != nil {
		log.Println("ChangeUserStatus FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

func UpdateUserRole(c *fiber.Ctx) error {
	rcol := DB.MI.RoleDBCol
	ucol := DB.MI.UserDBCol

	data := new(Model.User)
	auth := new(Model.Auth)
	rdata := new(Model.Roles)

	name := c.Query("name")
	role := c.Query("role")

	err := ucol.FindOne(c.Context(), bson.D{{"username", name}}).Decode(&data)
	if err != nil {
		log.Println("UpdateUserRole user FindOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = rcol.FindOne(c.Context(), bson.D{{"name", role}}).Decode(&rdata)
	if err != nil {
		log.Println("UpdateUserRole role FindOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	auth.Admin = rdata.Admin
	auth.Dashboard = rdata.Dashboard
	auth.Livechat = rdata.Livechat
	auth.Contact = rdata.Contact
	auth.Flowbuilder = rdata.Flowbuilder
	auth.Integrations = rdata.Integrations
	auth.Organization = rdata.Organization
	auth.ProductCatalogue = rdata.ProductCatalogue
	auth.Boardcast = rdata.Boardcast

	update := bson.M{"$set": bson.M{"role": role, "authority": auth}}

	_, err = ucol.UpdateOne(c.Context(), bson.D{{Key: "username", Value: name}}, update)
	if err != nil {
		log.Println("UpdateUserRole UpdateOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = ucol.FindOne(c.Context(), bson.D{{Key: "username", Value: name}}).Decode(&data)
	if err != nil {
		log.Println("UpdateUserRole FindOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

func UpdateDivisionAndTeam(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)
	if err := c.BodyParser(&user); err != nil {
		log.Println("UpdateDivisionAndTeam parser", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	update := bson.M{"$set": bson.M{"division_name": user.DivisionName, "team": user.Team}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		log.Println("UpdateDivisionAndTeam UpdateOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		log.Println("UpdateDivisionAndTeam FindOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteUserTeam(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)
	if err := c.BodyParser(&user); err != nil {
		log.Println("DeleteUserTeam parse", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	update := bson.M{"$set": bson.M{"division_name": "", "team": ""}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		log.Println("DeleteUserTeam UpdateOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		log.Println("DeleteUserTeam FindOne", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data)
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
	return c.Status(fiber.StatusCreated).JSON(chanInfo)
}
