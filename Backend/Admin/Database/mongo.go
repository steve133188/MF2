package Database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)



func MongoConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		fmt.Println("Cannot connect database")
	}
	users := client.Database("logs").Collection("logs")

	counts , err := users.CountDocuments(context.TODO(),bson.D{} )
	if counts == 0 {
		res , err := users.InsertOne(context.TODO() , bson.M{"name":"steve" , "id":"1"})
		if err !=  nil{
			fmt.Println(err)
		}
		fmt.Println(res.InsertedID," Added")
	}

	defer cancel()
}




