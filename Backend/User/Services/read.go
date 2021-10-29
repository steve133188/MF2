package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	user := new(Model.User)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	paramID := c.Params("id")

	filter := bson.M{"id": paramID}

	findResult := usersCollection.FindOne(c.Context(), filter)
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	err := findResult.Decode(&user)
	// user.Date = user.Date.Add(time.Hour * 8)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})

}

func GetAllUsers(c *fiber.Ctx) error {
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
		"data": fiber.Map{
			"users": users,
		},
	})
}

func GetUsersById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.User{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	// customer.Date = customer.Date.Add(time.Hour * 8)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}

func GetUserByUsername(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	paramID := c.Params("username")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.User{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	// customer.Date = customer.Date.Add(time.Hour * 8)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}

func GetUserByEmail(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	paramID := c.Params("email")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.User{}

	query := bson.D{{Key: "email", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   err,
		})
	}

	// customer.Date = customer.Date.Add(time.Hour * 8)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"customer": customer,
		},
	})
}

func GetValidPassword(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	paramPassword := c.Params("password")
	paramEmail := c.Params("email")
	user := &Model.User{}

	query := bson.D{{Key: "email", Value: paramEmail}}

	err := collection.FindOne(c.Context(), query).Decode(user)
	if err != nil {
		fmt.Println(err)
	}

	match := Util.CheckPasswordHash(paramPassword, user.Password)

	if !match {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
		})
	}

	// customer.Date = customer.Date.Add(time.Hour * 8)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}
