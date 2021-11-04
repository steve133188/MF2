package DB

import (
	"context"
	"fmt"
	"mf-chat-services/Util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DBCol  *mongo.Collection
}

var MI MongoInstance

func MongoConnect() {
	url, name, c := Util.GoDotEnvVariable("DB_URL"), Util.GoDotEnvVariable("DB_NAME"), Util.GoDotEnvVariable("DB_COLLECTION")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	collection := client.Database(name).Collection(c)

	fmt.Println("DB connected!")
	MI = MongoInstance{
		Client: client,
		DBCol:  collection,
	}
}
