package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(c *fiber.Ctx) error {
	fmt.Println("getall")

	usersCollection := DB.MI.DBCol
	// Query to filter
	query := bson.D{{}}

	cursor, err := usersCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var users []Model.User = make([]Model.User, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	// for i := range users {
	// 	users[i].Date = users[i].Date.Add(time.Hour * 8)
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func GetUsersByTeam(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol
	data := new(Model.Div)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	query := bson.D{{"division_name", data.Division}, {"team", data.Team}}

	cursor, err := customerCollection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	var users []Model.User = make([]Model.User, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func GetUserByPhone(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol
	data := new(Model.Param)
	user := new(Model.User)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	query := bson.D{{"phone", data.Param}}

	err = customerCollection.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func GetUserAuthority(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol
	var data struct {
		Param []string `json:"param" bson:"param"`
	}
	user := new(Model.User)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	query := bson.D{{"authority", bson.D{{"$in", data.Param}}}}

	err = customerCollection.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func GetUserList(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var data []struct {
		UserName string `json:"username" bson:"username"`
	}

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	err = cursor.All(c.Context(), &data)
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range data {
		if v.UserName != "" {
			name = append(name, v.UserName)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    name,
	})
}

func GetUserByEmail(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol
	data := new(Model.Param)
	user := new(Model.User)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	query := bson.D{{"email", data.Param}}

	err = customerCollection.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func GetUserByUsername(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol
	data := new(Model.Param)
	user := new(Model.User)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	query := bson.D{{"username", data.Param}}

	err = customerCollection.FindOne(c.Context(), query).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
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
