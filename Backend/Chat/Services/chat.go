package Services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mf-chat-services/DB"
	"mf-chat-services/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChatRoom(room string, user string) error {
	log.Println("=====================CreateChatRoom=========================")
	col := DB.MI.DBCol

	chat := new(Model.Chat)

	chat.RoomID = make([]string, 0)
	chat.RoomID = append(chat.RoomID, room)
	chat.UserID = user

	filter := bson.D{{"user_id", user}}
	update := bson.D{{"$push", bson.D{{"room_id", room}}}}
	res := col.FindOneAndUpdate(context.TODO(), filter, update)
	if res.Err() == mongo.ErrNoDocuments {
		res, err := col.InsertOne(context.Background(), chat)
		if err != nil {
			log.Println("Failed to insert new data       ", err)
			return err
		}
		fmt.Printf("res: %v\n", res)
	}
	return nil
}

func DeleteChatRoom(room string, user string) error {
	log.Println("=====================DeleteChatRoom=========================")

	col := DB.MI.DBCol

	filter := bson.D{{"user_id", user}}
	update := bson.D{{"$pull", bson.D{{"room_id", room}}}}
	res, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("DeleteChatRoom UpdateOne failed       ", err)
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("failed to delete roomid from user")
	}

	return nil
}
