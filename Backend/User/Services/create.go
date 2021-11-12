package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddManyAgent(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol

	// var datas []Model.User = make([]Model.User, 0)
	type data []interface{}
	var datas data
	err := c.BodyParser(&datas)
	if err != nil {
		log.Println("AddManyAgent parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// err := json.Unmarshal(c.Body(), &datas)
	_, err = usersCollection.InsertMany(c.Context(), datas)
	if err != nil {
		log.Println("AddManyAgent InsertMany: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func AddAgent(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol

	data := new(Model.User)
	exist := new(Model.User)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddAgent parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	data.CreatedAt = time.Now().Format("January 2 2006 15:04:05")
	data.Password, err = Util.HashPassword(data.Password)
	if err != nil {
		log.Println("AddAgent HashPassword: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	emailExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "email", Value: data.Email}}).Decode(exist)
	if (emailExisted) == nil {
		log.Println("AddAgent FindOne email: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userNameExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: data.UserName}}).Decode(exist)
	if userNameExisted == nil {
		log.Println("AddAgent FindOne name: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = usersCollection.InsertOne(c.Context(), data)
	if err != nil {
		log.Println("AddAgent InsertOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// get the inserted data
	user := &Model.User{}
	query := bson.D{{"phone", data.Phone}}
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(user)
}

// func TestLogin(c *fiber.Ctx) error {
// 	token := c.Request().Header.Peek("Authorization")
// 	_, err := Util.ParseToken(string(token))
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"success": false,
// 		})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 	})
// }
