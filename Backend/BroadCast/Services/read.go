package Services

import (
	"log"
	"mf-broadCast-services/DB"
	"mf-broadCast-services/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBroadCasts(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		log.Println("GetAllBroadCasts Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		log.Println("GetAllBroadCasts All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetBroadCastsByGroup(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Param)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("GetBroadCastsByGroup parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	query := bson.D{{"group", data.Param}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		log.Println("GetBroadCastsByGroup find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	err = cursor.All(c.Context(), &todos)
	if err != nil {
		log.Println("GetBroadCastsByGroup All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetBroadCastsByName(c *fiber.Ctx) error {
	collection := DB.MI.DBCol

	data := new(Model.Param)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("GetBroadCastsByName parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	query := bson.D{{"name", data.Param}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		log.Println("GetBroadCastsByName Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var todos []Model.BroadCast = make([]Model.BroadCast, 0)

	err = cursor.All(c.Context(), &todos)
	if err != nil {
		log.Println("GetBroadCastsByName all ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}
