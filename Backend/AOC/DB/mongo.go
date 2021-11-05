package DB

import (
	"context"
	"fmt"

	"mf-aoc-service/Util"

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
	RoleDBCol  *mongo.Collection
	ChanClient *mongo.Client
	ChanDBCol  *mongo.Collection
	TagsClient *mongo.Client
	TagsDBCol  *mongo.Collection
	OrgClient  *mongo.Client
	OrgDBCol   *mongo.Collection
	GrpClient  *mongo.Client
	GrpDBCol   *mongo.Collection
}

var MI MongoInstance

func MongoConnect() {
	chanUrl, channelDB, chanCol := Util.GoDotEnvVariable("CHAN_URL"), Util.GoDotEnvVariable("CHAN_NAME"), Util.GoDotEnvVariable("CHAN_COLLECTION")
	adminUrl, adminDB, roleCol, tagsCol := Util.GoDotEnvVariable("ADMIN_URL"), Util.GoDotEnvVariable("ADMIN_NAME"), Util.GoDotEnvVariable("ROLE_COLLECTION"), Util.GoDotEnvVariable("TAGS_COLLECTION")
	orgUrl, orgDB, orgCol := Util.GoDotEnvVariable("ORG_URL"), Util.GoDotEnvVariable("ORG_NAME"), Util.GoDotEnvVariable("ORG_COLLECTION")
	grpCol := Util.GoDotEnvVariable("GRP_COLLECTION")

	// adminCol := Util.GoDotEnvVariable("ADMIN_COLLECTION")
	// orgCol := Util.GoDotEnvVariable("ORG_COLLECTION")

	ctx := context.Background()
	grp, err := mongo.Connect(ctx, options.Client().ApplyURI(adminUrl))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	grps := grp.Database(adminDB).Collection(grpCol)

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

	// }

	fmt.Println("DB connected!")
	MI = MongoInstance{
		GrpClient:  grp,
		GrpDBCol:   grps,
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
