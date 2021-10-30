package Services

import (
	"fmt"
	"mf-channel-service/DB"
	"mf-channel-service/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllChannelInfo(c *fiber.Ctx) error {
	fmt.Println("getall")
	// token := c.Request().Header.Peek("Authorization")
	// _, err := Util.ParseToken(string(token))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Unauthorized",
	// 	})
	// }

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

	var users []Model.Channel = make([]Model.Channel, 0)

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

func GetChannelInfoById(c *fiber.Ctx) error {
	customerCollection := DB.MI.DBCol

	// get parameter value
	// paramID, err := primitive.ObjectIDFromHex(c.Params("id")) //valid id: 24 hex
	// if err != nil {
	// 	fmt.Println(err)
	// }

	paramID := c.Params("id")
	fmt.Println(paramID)

	// find todo and return
	customer := &Model.Channel{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := customerCollection.FindOne(c.Context(), query).Decode(customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	// timezone, err := strconv.ParseInt(customer.TimeZone, 10, 64)
	// if err != nil {
	// 	fmt.Println("Failed to parse TimeZone to int64, using default Time Zone UTC +8")
	// 	timezone = 8
	// }

	// customer.AccountCreatedTime = customer.AccountCreatedTime.Add(time.Hour * time.Duration(timezone))
	// customer.LastUpdatedTime = customer.LastUpdatedTime.Add(time.Hour * time.Duration(timezone))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}
