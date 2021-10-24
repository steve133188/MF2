package DB

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	url        string
	name       string
	collection string
	user       string
	pwd        string
}

type MongoInstance struct {
	Client *mongo.Client
	DBCol  *mongo.Collection
}

var MI MongoInstance

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func MongoConnect() {
	url, name, c := goDotEnvVariable("DB_URL"), goDotEnvVariable("DB_NAME"), goDotEnvVariable("DB_COLLECTION")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	collection := client.Database(name).Collection(c)
	count, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	// counts, err := collection.CountDocuments(context.TODO(), bson.D{})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if counts > 0 {
	// 	res, err := collection.DeleteMany(ctx, bson.M{})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(res.DeletedCount, " Deleted")
	// 	counts = 0
	// }

	// if counts == 0 {
	// 	res, err := collection.InsertOne(context.TODO(), bson.M{"id": "1", "userId": "111", "customerFirstName": "Tom", "customerLastName": "Boy", "age": "20", "date": time.Now()})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(res.InsertedID, " Added")
	// }

	fmt.Println("DB connected!")
	MI = MongoInstance{
		Client: client,
		DBCol:  collection,
	}
}
