package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ForgotPassword(c *fiber.Ctx) error {
	col := DB.MI.DBCol
	var target struct {
		Address string `json:"address"`
	}

	err := c.BodyParser(&target)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse",
			"error":   err.Error(),
		})
	}
	user := new(Model.User)
	err = col.FindOne(c.Context(), bson.D{{"email", target.Address}}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}
	randomPassword := Util.GeneratePassword(2, 2, 2, 8)
	password, err := Util.HashPassword(randomPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}
	result, err := col.UpdateOne(c.Context(), bson.D{{"email", target.Address}}, bson.D{{"$set", bson.D{{"password", password}}}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}

	err = Util.SendEmail(target.Address, randomPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to send mail",
			"error":   err.Error(),
		})
	}
	// Receiver email address.
	return c.Status(fiber.StatusOK).JSON(result)
}

func Login(c *fiber.Ctx) error {
	fmt.Println("Login")
	collection := DB.MI.DBCol
	// paramPassword := c.Params("password")
	// paramEmail := c.Params("email")
	user := new(Model.User)
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

	find.LastLogin = time.Now().Format("January 2 2006 15:04:05")
	_, err = collection.UpdateOne(c.Context(), bson.D{{"email", user.Email}}, bson.D{{"$set", bson.D{{"last_login", find.LastLogin}}}})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err.Error(),
		})
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

// func ChangeUserPassword(c *fiber.Ctx) error {
// 	usersCollection := DB.MI.DBCol
// 	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	user := new(Model.User)

// 	if err := c.BodyParser(&user); err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Failed to parse body",
// 			"error":   err.Error(),
// 		})
// 	}
// 	password, err := Util.HashPassword(user.Password)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot insert agent",
// 			"error":   err.Error(),
// 		})
// 	}
// 	update := bson.M{"$set": bson.M{"password": password}}

// 	err = usersCollection.FindOneAndUpdate(c.Context(), bson.D{{Key: "email", Value: user.Email}}, update).Decode(&user)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "User failed to update",
// 			"error":   err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 	})
// }
