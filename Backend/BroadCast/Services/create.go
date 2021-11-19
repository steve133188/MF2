package Services

import (
	"encoding/json"
	"fmt"
	"log"
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddBroadCast(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.BroadCast)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddBroadCast parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	data.CreatedDate = time.Now().Format("January 2 2006 15:04:05")

	_, err = collection.InsertOne(c.Context(), data)
	if err != nil {
		log.Println("AddBroadCast InsertOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// get the inserted data
	todo := &Model.BroadCast{}
	query := bson.D{{"name", data.Name}}

	collection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func AddManyBroadCast(c *fiber.Ctx) error {
	col := DB.MI.DBCol

	var datas []interface{}
	todo := new(Model.BroadCast)

	err := c.BodyParser(&datas)
	if err != nil {
		log.Println("AddManyBroadCast parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
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

		todo.CreatedDate = time.Now().Format("January 2 2006 15:04:05")

		_, err = col.InsertOne(c.Context(), todo)
		if err != nil {
			log.Println("AddManyBroadCast InsertOne ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

	}

	cursor, err := col.Find(c.Context(), bson.D{{}})
	if err != nil {
		log.Println("AddManyBroadCast Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		log.Println("AddManyBroadCast All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}
