package DB

import (
	"context"
	"fmt"

	"mf-aoc-service/Util"

	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type DB struct {
// 	url        string
// 	name       string
// 	collection string
// 	user       string
// 	pwd        string
// }

type MongoInstance struct {
	AdminClient *mongo.Client
	ChanClient  *mongo.Client
	OrgClient   *mongo.Client
	AdminDBCol  *mongo.Collection
	ChanDBCol   *mongo.Collection
	OrgDBCol    *mongo.Collection
}

var MI MongoInstance

func MongoConnect() {
	adminCol := Util.GoDotEnvVariable("ADMIN_DB_COLLECTION")
	chanUrl, channelDB, chanCol := Util.GoDotEnvVariable("CHAN_DB_URL"), Util.GoDotEnvVariable("CHAN_DB_NAME"), Util.GoDotEnvVariable("CHAN_DB_COLLECTION")
	orgCol := Util.GoDotEnvVariable("ORG_DB_COLLECTION")

	ctx := context.Background()
	admin, err := mongo.Connect(ctx, options.Client().ApplyURI(chanUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	admins := admin.Database(channelDB).Collection(adminCol)

	channel, err := mongo.Connect(ctx, options.Client().ApplyURI(chanUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	channels := channel.Database(channelDB).Collection(chanCol)

	org, err := mongo.Connect(ctx, options.Client().ApplyURI(chanUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	orgs := org.Database(channelDB).Collection(orgCol)

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Failed to generate first uuid")
	}
	res, err := channels.InsertOne(context.TODO(), bson.M{"id": id.String(), "name": "whatsappstella", "channel_id": "5e4367dd3c660d5d5e541176", "server": " https://35.198.244.95:9099", "server_auth": Util.GoDotEnvVariable("Stella_token")})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID, " Added")
	// }

	fmt.Println("DB connected!")
	MI = MongoInstance{
		AdminClient: admin,
		AdminDBCol:  admins,
		ChanClient:  channel,
		ChanDBCol:   channels,
		OrgClient:   org,
		OrgDBCol:    orgs,
	}
}
