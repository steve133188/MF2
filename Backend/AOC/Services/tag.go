package Services

import (
	"fmt"
	"log"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Post
func AddTags(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	data := new(Model.Tags)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()
	data.ID = id.String()
	data.Created = time.Now().Format("January 2 2006 15:04:05")
	data.Updated = time.Now().Format("January 2 2006 15:04:05")

	result, err := collection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	// get the inserted data
	todo := &Model.Role{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

//Delete
func DeleteTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	// get param
	paramID := c.Params("name")

	// find and delete todo
	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

//Get
func GetAllTags(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	// Query to filter
	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}
	var todos []Model.Tags = make([]Model.Tags, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	paramID := c.Params("name")
	fmt.Println(paramID)

	// find todo and return
	todo := &Model.Tags{}

	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)

	// todo.Date = todo.Date.Add(time.Hour * 8)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

//Update
func UpdateTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Tags)

	if err := c.BodyParser(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "name", Value: c.Params("name")}}, update)
	fmt.Println(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}
