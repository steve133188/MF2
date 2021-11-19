package Services

import (
	"fmt"
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ForgotPassword(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol
	var target struct {
		Address string `json:"address"`
	}

	err := c.BodyParser(&target)
	if err != nil {
		log.Println("ForgotPassword parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user := new(Model.User)
	err = col.FindOne(c.Context(), bson.D{{"email", target.Address}}).Decode(&user)
	if err != nil {
		log.Println("ForgotPassword FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	randomPassword := Util.GeneratePassword(2, 2, 2, 8)
	password, err := Util.HashPassword(randomPassword)
	if err != nil {
		log.Println("ForgotPassword HashPassword: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	_, err = col.UpdateOne(c.Context(), bson.D{{"email", target.Address}}, bson.D{{"$set", bson.D{{"password", password}}}})
	if err != nil {
		log.Println("ForgotPassword UpdateOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = Util.SendEmail(target.Address, randomPassword)
	if err != nil {
		log.Println("ForgotPassword SendEmail: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Receiver email address.
	return c.SendStatus(fiber.StatusOK)
}

func Login(c *fiber.Ctx) error {
	fmt.Println("Login")
	collection := DB.MI.UserDBCol
	// paramPassword := c.Params("password")
	// paramEmail := c.Params("email")
	user := new(Model.User)
	find := new(Model.User)

	err := c.BodyParser(&user)
	if err != nil {
		log.Println("Login parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	query := bson.D{{Key: "email", Value: user.Email}}

	err = collection.FindOne(c.Context(), query).Decode(find)
	if err != nil {
		log.Println("Login FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	match := Util.CheckPasswordHash(user.Password, find.Password)
	if !match {
		log.Println("Login CheckPasswordHash: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	find.LastLogin = time.Now().Format("January 2 2006 15:04:05")
	_, err = collection.UpdateOne(c.Context(), bson.D{{"email", user.Email}}, bson.D{{"$set", bson.D{{"last_login", find.LastLogin}}}})
	if err != nil {
		log.Println("Login UpdateOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["username"] = find.UserName
	claims["password"] = find.Password
	claims["email"] = find.Email
	claims["role"] = find.Role
	claims["status"] = find.Status

	Secret := Util.GoDotEnvVariable("Token_pwd")
	s, err := token.SignedString([]byte(Secret))
	if err != nil {
		fmt.Println(err)
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodPS256, tk)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    find,
		"token":   s,
	})

}

func ChangeUserPassword(c *fiber.Ctx) error {
	usersCollection := DB.MI.UserDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}
	result := new(Model.User)

	if err := c.BodyParser(&user); err != nil {
		log.Println("ChangeUserPassword parse", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	password, err := Util.HashPassword(user.Password)
	if err != nil {
		log.Println("ChangeUserPassword hashpassword ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	update := bson.M{"$set": bson.M{"password": password}}

	err = usersCollection.FindOneAndUpdate(c.Context(), bson.D{{Key: "email", Value: user.Email}}, update).Decode(&result)
	if err != nil {
		log.Println("ChangeUserPassword FindOneAndUpdate", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
