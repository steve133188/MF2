package Services

import (
	"log"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"

	"github.com/rs/xid"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetReplyFolderByID(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	paramID := c.Params("id")

	reply := &Model.StandardReply{}
	filter := bson.D{{Key: "id", Value: paramID}}

	err := adminColl.FindOne(c.Context(), filter).Decode(reply)
	if err != nil {
		log.Println("GetReplyFolderByID FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(reply)
}

func GetAllReplyFolder(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	cursor, err := adminColl.Find(c.Context(), bson.D{{}})
	if err != nil {
		log.Println("GetAllReplyFolder Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var replies []Model.StandardReply = make([]Model.StandardReply, 0)
	err = cursor.All(c.Context(), &replies)
	if err != nil {
		log.Println("GetAllReplyFolder All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(replies)
}

func CreateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		log.Println("CreateReply parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	id := xid.New()
	reply.ID = id.String()
	reply.Content = make([]Model.Content, 0)

	sameNameFilter := bson.D{{Key: "name", Value: reply.Name}}
	count, err := adminColl.CountDocuments(c.Context(), sameNameFilter)
	if count > 0 {
		log.Println("CreateReply CountDocuments ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	result, err := adminColl.InsertOne(c.Context(), reply)
	if err != nil {
		log.Println("CreateReply InsertOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//check the inserted data and return
	checkReply := &Model.StandardReply{}
	checkFilter := bson.D{{Key: "_id", Value: result.InsertedID}}

	err = adminColl.FindOne(c.Context(), checkFilter).Decode(checkReply)
	if err != nil {
		log.Println("CreateReply FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(checkReply)
}

func UpdateReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	reply := new(Model.StandardReply)

	err := c.BodyParser(&reply)
	if err != nil {
		log.Println("UpdateReply parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	target := bson.D{{Key: "id", Value: reply.ID}}
	update := bson.D{{"$set", reply}}

	result, err := adminColl.UpdateOne(c.Context(), target, update)
	if err != nil {
		log.Println("UpdateReply UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	checkReply := new(Model.StandardReply)
	checkFilter := bson.D{{Key: "_id", Value: result.UpsertedID}}

	err = adminColl.FindOne(c.Context(), checkFilter).Decode(&checkReply)
	if err != nil {
		log.Println("UpdateReply FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(checkReply)
}

func DeleteReply(c *fiber.Ctx) error {
	adminColl := DB.MI.RpyDBCol

	paramID := c.Params("id")
	filter := bson.D{{Key: "id", Value: paramID}}

	result, err := adminColl.DeleteOne(c.Context(), filter)
	if err != nil {
		log.Println("DeleteReply DeleteOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result.DeletedCount)
}

func AddContent(c *fiber.Ctx) error {
	col := DB.MI.RpyDBCol
	rpyContent := new(Model.StandardReply)
	content := new(Model.Content)
	recContent := new(Model.MessageContent)
	err := c.BodyParser(&recContent)
	if err != nil {
		log.Println("AddContent parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	folder := recContent.FolderName
	content.ID = xid.New().String()
	content.Body = recContent.ContentBody

	update := bson.D{{"$push", bson.D{{"content", content}}}}
	filter := bson.D{{"name", folder}}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("AddContent UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Println(res.ModifiedCount)

	err = col.FindOne(c.Context(), filter).Decode(&rpyContent)
	if err != nil {
		log.Println("AddContent FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(rpyContent.Content)
}

func UpdateContent(c *fiber.Ctx) error {
	col := DB.MI.RpyDBCol

	recContent := new(Model.Content)

	err := c.BodyParser(&recContent)
	if err != nil {
		log.Println("UpdateContent parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"content.id", recContent.ID}}
	update := bson.D{{"$set", bson.D{{"content.$.body", recContent.Body}}}}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("UpdateContent UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("UpdateContent count ", err)
		return c.SendStatus(fiber.StatusNotModified)
	}

	rpyContent := new(Model.StandardReply)
	err = col.FindOne(c.Context(), filter).Decode(&rpyContent)
	if err != nil {
		log.Println("UpdateContent FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(rpyContent)
}

func DeleteContent(c *fiber.Ctx) error {
	col := DB.MI.RpyDBCol

	recContent := new(Model.MessageContent)
	content := new(Model.Content)

	err := c.BodyParser(&recContent)
	if err != nil {
		log.Println("DeleteContent parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	content.Body = recContent.ContentBody
	content.ID = recContent.ContentID

	filter := bson.D{
		{"name", recContent.FolderName},
	}

	update := bson.D{{"$pull", bson.D{{"content", content}}}}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("DeleteContent UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("DeleteContent count ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	msgContent := new(Model.StandardReply)
	err = col.FindOne(c.Context(), filter).Decode(&msgContent)
	if err != nil {
		log.Println("DeleteContent FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(msgContent)
}

func GetContentsByFolderName(c *fiber.Ctx) error {
	col := DB.MI.RpyDBCol

	folderName := c.Params("name")

	rpyContent := new(Model.StandardReply)

	filter := bson.D{{"name", folderName}}
	err := col.FindOne(c.Context(), filter).Decode(&rpyContent)
	if err != nil {
		log.Println("GetContentsByFolderName FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(rpyContent.Content)

}
