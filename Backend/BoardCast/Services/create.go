package Services

import (
	"encoding/json"
	"fmt"
	"mf-boardCast-services/DB"
	"mf-boardCast-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddBoardCast(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.BoardCast)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	data.CreatedDate = time.Now().Format("Jan 2, 2006")

	_, err = collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert customer",
			"error":   err.Error(),
		})
	}

	// get the inserted data
	todo := &Model.BoardCast{}
	query := bson.D{{"name", data.Name}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func AddManyBoardCast(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var datas []interface{}
	todo := new(Model.BoardCast)

	err := c.BodyParser(&datas)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	for _, v := range datas {
		mapData := v.(map[string]interface{})
		b, err := json.Marshal(mapData)
		if err != nil {
			fmt.Println("send response marshal error")
		}
		err = json.Unmarshal(b, &todo)
		if err != nil {
			fmt.Println("send response unmarshal error")
		}

		todo.CreatedDate = time.Now().Format("Jan 2, 2006")

		_, err = col.InsertOne(c.Context(), todo)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot insert customer",
				"error":   err.Error(),
			})
		}

	}

	cursor, err := col.Find(c.Context(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.BoardCast = make([]Model.BoardCast, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todos,
	})
}
