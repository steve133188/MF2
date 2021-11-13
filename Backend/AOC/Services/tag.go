package Services

import (
	"log"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Post
func AddTag(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	data := new(Model.Tags)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddTag parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	id := xid.New()
	data.ID = id.String()
	data.Created = time.Now().Format("January 2 2006 15:04:05")
	data.Updated = time.Now().Format("January 2 2006 15:04:05")

	result, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		log.Println("AddTag InsertOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// get the inserted data
	todo := &Model.Tags{}
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
			if err != nil {
				log.Println("DeleteTagsByName ", err)
				return c.SendStatus(fiber.StatusNotFound)
			}
		}

		log.Println("DeleteTagsByName ", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

//Get
func GetAllTags(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	// Query to filter
	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		log.Println("GetAllTags find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var todos []Model.Tags = make([]Model.Tags, 0)
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		log.Println("GetAllTags All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTagByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	paramID := c.Params("name")

	todo := &Model.Tags{}

	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)
	if err != nil {
		log.Println("GetTagsByName FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

//Update
func UpdateTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Tags)
	if err := c.BodyParser(todo); err != nil {
		log.Println("UpdateTagsByName parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "name", Value: c.Params("name")}}, update)
	if err != nil {
		log.Println("UpdateTagsByName UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetTagList(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol

	var data []struct {
		Tags string `json:"tags"`
	}

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("GetTagList Find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &data)
	if err != nil {
		log.Println("GetTagList All: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range data {
		if v.Tags != "" {
			name = append(name, v.Tags)
		}
	}

	return c.Status(fiber.StatusOK).JSON(name)
}
