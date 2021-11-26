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

	count, err := usersCollection.CountDocuments(c.Context(), bson.D{{"email", data.Email}})
	if err != nil {
		log.Println("AddAgent count email: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if count != 0 {
		log.Println("email existed: ", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	count, err = usersCollection.CountDocuments(c.Context(), bson.D{{"username", data.UserName}})
	if err != nil {
		log.Println("AddAgent count name: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if count != 0 {
		log.Println("name existed: ", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	count, err = usersCollection.CountDocuments(c.Context(), bson.D{{"phone", data.Phone}})
	if err != nil {
		log.Println("AddAgent count phone: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if count != 0 {
		log.Println("phone existed: ", err)
		return c.SendStatus(fiber.StatusBadRequest)
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
