package DB

import (
	"context"
	"fmt"
	"mf-user-servies/Util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	UserClient *mongo.Client
	UserDBCol  *mongo.Collection
}

var MI MongoInstance

func MongoConnect() {
	ctx := context.Background()

	user, err := mongo.Connect(ctx, options.Client().ApplyURI(Util.GoDotEnvVariable("DB_URL")))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	users := user.Database(Util.GoDotEnvVariable("DB_NAME")).Collection(Util.GoDotEnvVariable("USER_COLLECTION"))

	fmt.Println("DB connected!")
	MI = MongoInstance{
		UserClient: user,
		UserDBCol:  users,
	}

}
