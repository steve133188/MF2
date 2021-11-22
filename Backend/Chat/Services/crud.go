package Services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mf-chat-services/DB"
	"mf-chat-services/Model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleChat(w http.ResponseWriter, r *http.Request) {
	col := DB.MI.DBCol
	vars := mux.Vars(r)
	userId := vars["id"]
	chat := new(Model.Chat)

	if r.Method == "GET" {
		query := bson.D{{"user_id", userId}}
		err := col.FindOne(r.Context(), query).Decode(&chat)
		if err != nil {
			log.Println("Get FindOne      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(chat)
		if err != nil {
			log.Println("HandleChat marshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else if r.Method == "DELETE" {
		filter := bson.D{{"user_id", userId}}
		todo, err := col.DeleteOne(r.Context(), filter)
		if err != nil {
			log.Println("Delete Delete One      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(todo)
		if err != nil {
			log.Println("HandleChat marshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	}
}

func HandleGetAllChat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		col := DB.MI.DBCol

		var chats []Model.Chat = make([]Model.Chat, 0)

		cursor, err := col.Find(r.Context(), bson.D{{}})
		if err != nil {
			log.Println("HandleGetAllChat find      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer cursor.Close(r.Context())

		err = cursor.All(r.Context(), &chats)
		if err != nil {
			log.Println("HandleGetAllChat All      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(chats)
		if err != nil {
			log.Println("HandleChat marshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func HandleUpdateOneChat(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("HandleUpdateOneChat ReadAll      ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	chat := new(Model.Chat)

	err = json.Unmarshal(data, &chat)
	if err != nil {
		log.Println("HandleUpdateOneChat unmarshal      ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method == "PUT" {

		filter := bson.D{{"user_id", chat.UserID}}
		update := bson.D{{"$set", chat}}

		col := DB.MI.DBCol
		res, err := col.UpdateOne(r.Context(), filter, update)
		if err != nil {
			log.Println("HandleUpdateOneChat UpdateOne      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if res.ModifiedCount == 0 {
			log.Println("Update failed      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = col.FindOne(r.Context(), filter).Decode(&chat)
		if err != nil {
			log.Println("HandleUpdateOneChat FindOne      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		val, err := json.Marshal(chat)
		if err != nil {
			log.Println("HandleChat marshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(val)
	} else if r.Method == "POST" {
		col := DB.MI.DBCol

		userId := chat.UserID
		count, err := col.CountDocuments(r.Context(), bson.D{{"user_id", userId}})
		if count > 0 {
			log.Println("Document existed      ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := col.InsertOne(r.Context(), chat)
		if err != nil {
			log.Println("HandleUpdateOneChat InsertOne      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		filter := bson.D{{"_id", res.InsertedID}}
		err = col.FindOne(r.Context(), filter).Decode(&chat)
		if err != nil {
			log.Println("HandleUpdateOneChat FindOne      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		val, err := json.Marshal(chat)
		if err != nil {
			log.Println("HandleChat marshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(val)
	}

}
