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

//func NewDB() *DB{
//	url := goDotEnvVariable("DB_URL")
//	name := goDotEnvVariable("DB_NAME")
//	c := goDotEnvVariable("DB_COLLECTION")
//	return &DB{url:url , name : name , collection: c}
//}

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
	users := client.Database(name).Collection(c) // name constant by service
	count, err := users.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	counts, err := users.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	//test data insert start

	if counts > 0 {
		res, err := users.DeleteMany(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.DeletedCount, " Deleted")
		counts = 0
	}
	if counts == 0 {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println("Failed to generate first ID")
		}
		res1, err := users.InsertOne(ctx, bson.M{"id": id.String(), "username": "steve", "password": "12345", "emails": "stevechakcy@gmail.com", "created_at": time.Now()})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res1.InsertedID, " Added")
	}

	//end of test data insert

	fmt.Println("DB connected!")
	MI = MongoInstance{
		Client: client,
		DBCol:  users,
	}

}
