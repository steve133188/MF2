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
	RoleClient *mongo.Client
	TagsClient *mongo.Client
	ChanClient *mongo.Client
	OrgClient  *mongo.Client
	RoleDBCol  *mongo.Collection
	TagsDBCol  *mongo.Collection
	ChanDBCol  *mongo.Collection
	OrgDBCol   *mongo.Collection
}

var MI MongoInstance

func MongoConnect() {
	chanUrl, channelDB, chanCol := Util.GoDotEnvVariable("CHAN_URL"), Util.GoDotEnvVariable("CHAN_NAME"), Util.GoDotEnvVariable("CHAN_COLLECTION")
	adminUrl, adminDB, roleCol, tagsCol := Util.GoDotEnvVariable("ADMIN_URL"), Util.GoDotEnvVariable("ADMIN_NAME"), Util.GoDotEnvVariable("ROLE_COLLECTION"), Util.GoDotEnvVariable("TAGS_COLLECTION")
	orgUrl, orgDB, orgCol := Util.GoDotEnvVariable("ORG_URL"), Util.GoDotEnvVariable("ORG_NAME"), Util.GoDotEnvVariable("ORG_COLLECTION")
	// adminCol := Util.GoDotEnvVariable("ADMIN_COLLECTION")
	// orgCol := Util.GoDotEnvVariable("ORG_COLLECTION")

	ctx := context.Background()
	role, err := mongo.Connect(ctx, options.Client().ApplyURI(adminUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	roles := role.Database(adminDB).Collection(roleCol)

	tag, err := mongo.Connect(ctx, options.Client().ApplyURI(adminUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	tags := tag.Database(adminDB).Collection(tagsCol)

	channel, err := mongo.Connect(ctx, options.Client().ApplyURI(chanUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	channels := channel.Database(channelDB).Collection(chanCol)

	org, err := mongo.Connect(ctx, options.Client().ApplyURI(orgUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	orgs := org.Database(orgDB).Collection(orgCol)

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
		RoleClient: role,
		RoleDBCol:  roles,
		TagsClient: tag,
		TagsDBCol:  tags,
		ChanClient: channel,
		ChanDBCol:  channels,
		OrgClient:  org,
		OrgDBCol:   orgs,
	}
}
