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

	data.CreatedAt = time.Now().Format("January 2, 2006 14:00")
	data.Password, err = Util.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert agent",
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

	find.LastLogin = time.Now().Format("January 2, 2006 14:00")
	_, err = collection.InsertOne(c.Context(), find)
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
