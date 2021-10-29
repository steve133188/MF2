package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	data := new(Model.User)
	exist := new(Model.User)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate ID in POST")
	}
	data.ID = id.String()
	data.CreatedAt = time.Now()
	data.Password, err = Util.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	emailExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "email", Value: data.Email}}).Decode(exist)
	if (emailExisted) == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User email exist",
		})
	}

	userNameExisted := usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: data.UserName}}).Decode(exist)
	fmt.Println(userNameExisted)
	if userNameExisted == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User name exist",
		})
	}

	result, err := usersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.User{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	fmt.Println(result.InsertedID)
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

func Login(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	// paramPassword := c.Params("password")
	// paramEmail := c.Params("email")
	user := &Model.User{}
	find := new(Model.User)

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	query := bson.D{{Key: "email", Value: user.Email}}

	err = collection.FindOne(c.Context(), query).Decode(find)
	if err != nil {
		fmt.Println(err)
	}

	match := Util.CheckPasswordHash(user.Password, find.Password)

	if !match {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":  false,
			"response": "failed to login",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7)
	claims["username"] = user.UserName
	claims["password"] = user.Password
	claims["email"] = user.Email

	Secret := Util.GoDotEnvVariable("Token_pwd")
	s, err := token.SignedString([]byte(Secret))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	// token := jwt.NewWithClaims(jwt.SigningMethodPS256, tk)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"response": "login success",
		"token":    s,
	})

}
