package DB

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/nu7hatch/gouuid"

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
	customers := client.Database(name).Collection(c)
	count, err := customers.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	counts, err := customers.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	if counts > 0 {
		res, err := customers.DeleteMany(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.DeletedCount, " Deleted")
		counts = 0
	}

	if counts == 0 {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println("Failed to generate first uuid")
		}
		res, err := customers.InsertOne(context.TODO(), bson.M{"id": id.String(), "userId": "111", "customerFirstName": "Tom", "customerLastName": "Boy", "age": "20", "timezone": "8", "lastUpdatedTime": time.Now(), "accountCreatedTime": time.Now()})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.InsertedID, " Added")
	}

	fmt.Println("DB connected!")
	MI = MongoInstance{
		Client: client,
		DBCol:  customers,
	}
}
