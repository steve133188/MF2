package Services

import (
	"context"
	"github.com/rs/xid"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetReplyFolderByID(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	paramID := c.Params("id")

	reply := &Model.StandardReply{}
	filter := bson.D{{Key : "id", Value : paramID}}

	err := adminColl.FindOne(c.Context(), filter).Decode(reply)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Log Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(reply)
}

func GetAllReplyFolder(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	cursor, err := adminColl.Find(c.Context(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Replies Not Found",
			"error":   err.Error(),
		})
	}

	var replies []Model.StandardReply = make([]Model.StandardReply, 0)
	err = cursor.All(c.Context(), &replies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}

	defer cursor.Close(context.TODO())

	return c.Status(fiber.StatusOK).JSON(replies)
}

func CreateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := xid.New()
	reply.ID = id.String()

	sameNameFilter := bson.D{{Key: "name", Value: reply.Name}}
	count, err := adminColl.CountDocuments(c.Context(), sameNameFilter)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Name Has Been Used",
		})
	}

	result, err := adminColl.InsertOne(c.Context(), reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert reply",
			"error":   err,
		})
	}

	//check the inserted data and return
	checkReply := &Model.StandardReply{}
	checkFilter := bson.D{{Key: "_id", Value: result.InsertedID}}

	adminColl.FindOne(c.Context(), checkFilter).Decode(checkReply)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"StandardReply": checkReply,
		},
	})
}

func UpdateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	target := bson.D{{Key: "id", Value: reply.ID}}
	update := bson.D{{"$set", reply}}

	result, err := adminColl.UpdateOne(c.Context(), target, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update reply",
			"error":   err,
		})
	}

	checkReply := new(Model.StandardReply)
	checkFilter := bson.D{{Key: "_id", Value: result.UpsertedID}}

	adminColl.FindOne(c.Context(), checkFilter).Decode(&checkReply)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"standardReply": checkReply,
		},
	})
}

func DeleteReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	paramID := c.Params("id")
	filter := bson.D{{Key : "id", Value : paramID}}

	result, err := adminColl.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Folder Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"number of deletion": result.DeletedCount,
	})
}