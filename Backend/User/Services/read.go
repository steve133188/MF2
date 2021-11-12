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

func GetTeamList(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	var data []struct {
		Team string `json:"team"`
	}

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("GetTeamList find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &data)
	if err != nil {
		log.Println("GetTeamList cursor all: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range data {
		if v.Team != "" {
			found := false
			for _, val := range name {
				if v.Team == val {
					found = true
					break
				}
			}
			if !found {
				name = append(name, v.Team)
			}
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

// func GetValidPassword(c *fiber.Ctx) error {
// 	collection := DB.MI.DBCol
// 	paramPassword := c.Params("password")
// 	paramEmail := c.Params("email")
// 	user := &Model.User{}

// 	query := bson.D{{Key: "email", Value: paramEmail}}

// 	err := collection.FindOne(c.Context(), query).Decode(user)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	match := Util.CheckPasswordHash(paramPassword, user.Password)

// 	if !match {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"success": false,
// 		})
// 	}

// 	// customer.Date = customer.Date.Add(time.Hour * 8)

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 	})

// }

func GetUsersByTeam(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	team := c.Query("team")
	division := c.Query("division")

	query := bson.D{{"division_name", division}, {"team", team}}

	cursor, err := col.Find(c.Context(), query)
	if err != nil {
		log.Println("GetUsersByTeam Error: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var users []Model.User = make([]Model.User, 0)

	err = cursor.All(c.Context(), &users)
	if err != nil {
		log.Println("GetUsersByTeam Error: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
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
