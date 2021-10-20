package DB

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)


type MongoInstance struct {
	Client *mongo.Client
	DBCol     *mongo.Collection
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
	url,name,c := goDotEnvVariable("DB_URL") ,goDotEnvVariable("DB_NAME") , goDotEnvVariable("DB_COLLECTION")

	ctx:= context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		fmt.Println("Cannot connect database")
	}
	users := client.Database(name).Collection(c) // name constant by service
	count , err := users.Find(context.TODO() , bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	counts , err := users.CountDocuments(context.TODO(),bson.D{} )

	//test data insert start

	if counts>0{
		res , err :=users.DeleteMany(ctx,bson.M{})
		if err !=  nil{
			fmt.Println(err)
		}
		fmt.Println(res.DeletedCount," Deleted")
		res1 , err := users.InsertOne(ctx, bson.M{"username":"steve" , "password":"1234" ,"email":"stevechakcy@gmail.com"})
		if err !=  nil{
			fmt.Println(err)
		}
		fmt.Println(res1.InsertedID," Added")
	}
	if counts == 0 {
		res1 , err := users.InsertOne(ctx, bson.M{"username":"steve" , "password":"12345" ,"email":"stevechakcy@gmail.com"})
		if err !=  nil{
			fmt.Println(err)
		}
		fmt.Println(res1.InsertedID," Added")
	}

	//end of test data insert

	fmt.Println("DB connected!")
	MI = MongoInstance{
		Client: client,
		DBCol: users,
	}

}


