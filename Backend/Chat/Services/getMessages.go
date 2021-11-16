package Services

import (
	"mf-chat-services/DB"
	"mf-chat-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllMessages(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Chat)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	// Query to filter
	query := bson.D{{"customer_id", data.CustomerID}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find message",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())
	var todos []Model.Chat = make([]Model.Chat, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into messages",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todos,
	})
}

// func GetOneMessageById(c *fiber.Ctx) error {
// 	collection := DB.MI.DBCol

// 	paramID := c.Params("id")
// 	fmt.Println(paramID)

// 	// find todo and return
// 	todo := &Model.Message{}

// 	query := bson.D{{Key: "id", Value: paramID}}

// 	err := collection.FindOne(c.Context(), query).Decode(todo)

// 	// timezone UTC +8
// 	// chat.Time = chat.Time.Add(time.Hour * 8)

// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Message Not found",
// 			"error":   err,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data": fiber.Map{
// 			"message": todo,
// 		},
// 	})
// }
