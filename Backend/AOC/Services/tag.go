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
func DeleteTagsByID(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			if err != nil {
				log.Println("DeleteTagsByID ", err)
				return c.SendStatus(fiber.StatusNotFound)
			}
		}

		log.Println("DeleteTagsByID ", err)
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

	query := bson.D{{Key: "tag", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)
	if err != nil {
		log.Println("GetTagsByName FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func GetTagByID(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	paramID := c.Params("id")

	todo := &Model.Tags{}

	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOne(c.Context(), query).Decode(todo)
	if err != nil {
		log.Println("GetTagByID FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

//Update
func UpdateTagsByID(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Tags)
	if err := c.BodyParser(&todo); err != nil {
		log.Println("UpdateTagsByID parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	todo.Updated = time.Now().Format("January 2 2006 15:04:05")
	update := bson.D{{"$set", todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{"id", todo.ID}}, update)
	if err != nil {
		log.Println("UpdateTagsByID UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	result := new(Model.Tags)

	err = collection.FindOne(c.Context(), bson.D{{"id", todo.ID}}).Decode((&result))
	if err != nil {
		log.Println("UpdateTagsByID FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetTagList(c *fiber.Ctx) error {
	col := DB.MI.TagsDBCol

	var tags []Model.Tags = make([]Model.Tags, 0)

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("GetTagList Find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &tags)
	if err != nil {
		log.Println("GetTagList All: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range tags {
		if v.Tag != "" {
			name = append(name, v.Tag)
		}
	}

	return c.Status(fiber.StatusOK).JSON(name)
}
