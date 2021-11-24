package Services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mf-bot-services/DB"
	"mf-bot-services/Model"
)

func GetOneBotMessageByID(c *fiber.Ctx) error {
	botMsgColl := DB.MI.DBCol

	paramID := c.Params("id")
	fmt.Println(paramID)

	msg := &Model.ChatBotReply{}

	filter := bson.D{{Key: "id", Value: paramID}}

	err := botMsgColl.FindOne(c.Context(), filter).Decode(msg)
	if err != nil {
		log.Println("GetReplyFolderByID FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(msg)
}

func CreateOneBotMessage(c *fiber.Ctx) error {
	botMsgColl := DB.MI.DBCol

	newMsg := new(Model.ChatBotReply)

	err := c.BodyParser(&newMsg)
	if err != nil {
		log.Println("CreateReply parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	id := xid.New()
	newMsg.ID = id.String()

	result, err := botMsgColl.InsertOne(c.Context(), newMsg)
	if err != nil {
		log.Println("CreateReply InsertOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// get the inserted data
	checkMsg := &Model.ChatBotReply{}
	checkFilter := bson.D{{Key: "_id", Value: result.InsertedID}}

	err = botMsgColl.FindOne(c.Context(), checkFilter).Decode(checkMsg)
	if err != nil {
		log.Println("CreateReply FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(checkMsg)
}

func DeleteOneBotMessageById(c *fiber.Ctx) error {

	// get param
	paramID := c.Params("id")

	//delete count
	var count *int

	recursiveDeleteMsg(c, paramID, count)
	if *count < 0 {
		log.Println("DeleteBotMsg Err, DeletedCount ", *count)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"DeleteCount": *count,
	})
}

func recursiveDeleteMsg(c *fiber.Ctx, msgID string, count *int) {
	botMsgColl := DB.MI.DBCol

	filter := bson.D{{Key: "id", Value: msgID}}

	deleteMsg := &Model.ChatBotReply{}
	err := botMsgColl.FindOneAndDelete(c.Context(), filter).Decode(deleteMsg)
	if err != nil {
		log.Println("DeleteBotMsg DeleteOne ", err)
		*count = *count * -1
		return
	} else {
		*count++
	}

	if len(deleteMsg.ChildrenID) == 0 {
		return
	} else {
		for i := 0; i < len(deleteMsg.ChildrenID); i++ {
			recursiveDeleteMsg(c, deleteMsg.ChildrenID[i], count)
		}
	}
}

func UpdateOneBotMessageById(c *fiber.Ctx) error {
	botMsgColl := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	msg := new(Model.ChatBotReply)

	err := c.BodyParser(&msg)
	if err != nil {
		log.Println("UpdateReply parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	target := bson.D{{Key: "id", Value: msg.ID}}
	update := bson.D{{"$set", msg}}

	result, err := botMsgColl.UpdateOne(c.Context(), target, update)
	if err != nil {
		log.Println("UpdateReply UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	checkMsg := new(Model.ChatBotReply)
	checkFilter := bson.D{{Key: "_id", Value: result.UpsertedID}}

	err = botMsgColl.FindOne(c.Context(), checkFilter).Decode(&checkMsg)
	if err != nil {
		log.Println("UpdateReply FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(checkMsg)
}
