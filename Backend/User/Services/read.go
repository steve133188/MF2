package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol

	query := bson.D{{}}

	cursor, err := usersCollection.Find(c.Context(), query)

	if err != nil {
		log.Println("GetAllUsers find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var users []Model.User = make([]Model.User, 0)

	err = cursor.All(c.Context(), &users)
	if err != nil {
		log.Println("GetAllUsers cursor all: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUserList(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var data []struct {
		UserName string `json:"username" bson:"username"`
	}

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("GetUserList find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &data)
	if err != nil {
		log.Println("GetUserList cursor all: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range data {
		if v.UserName != "" {
			name = append(name, v.UserName)
		}
	}

	return c.Status(fiber.StatusOK).JSON(name)
}

func GetUserByEmail(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol
	user := new(Model.User)
	email := c.Params("email")
	query := bson.D{{"email", email}}

	err := col.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		log.Println("GetUsersByEmail findone: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUserByName(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)
	name := c.Params("name")
	query := bson.D{{"username", name}}

	err := col.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		log.Println("GetUsersByName find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUserByPhone(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)
	phone := c.Params("phone")
	query := bson.D{{"phone", phone}}

	err := col.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		if err != nil {
			log.Println("GetUsersByPhone Error: ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
