package Services

import (
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddManyAgent(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	// var datas []Model.User = make([]Model.User, 0)
	type data []interface{}
	var datas data
	err := c.BodyParser(&datas)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	// err := json.Unmarshal(c.Body(), &datas)
	_, err = usersCollection.InsertMany(c.Context(), datas)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert agent",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

func AddAgent(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	data := new(Model.User)
	exist := new(Model.User)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	data.CreatedAt = time.Now().Format("January 2 2006 15:04:05")
	data.Password, err = Util.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}

	emailExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "email", Value: data.Email}}).Decode(exist)
	if (emailExisted) == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Agent email exist",
		})
	}

	userNameExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: data.UserName}}).Decode(exist)
	if userNameExisted == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Agent name exist",
		})
	}

	_, err = usersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.User{}
	query := bson.D{{"phone", data.Phone}}
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
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
