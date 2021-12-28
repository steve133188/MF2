package handler

import (
	"aws-lambda-customer/model"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func FieldHandler(Agents []int, TeamID int, Tags []int, dynaClient *dynamodb.Client) ([]model.User, model.Team, []model.Tag, error) {
	userID := Agents

	var users []model.User = make([]model.User, 0)

	for _, v := range userID {
		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("USERTABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Printf("FailedToGetUser, userID = %v, %s", v, err)
			return nil, model.Team{}, nil, err
		}
		user := new(model.User)

		err = attributevalue.UnmarshalMap(out.Item, &user)
		if err != nil {
			log.Printf("UnmarshalMapError, %s", err)
			return nil, model.Team{}, nil, err
		}

		users = append(users, *user)
	}
	teamID := TeamID

	team := new(model.Team)

	tout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("ORGTABLE")),
		Key: map[string]types.AttributeValue{
			"team_id": &types.AttributeValueMemberN{Value: strconv.Itoa(teamID)},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetUser, teamID = %v, %s", teamID, err)
		return nil, model.Team{}, nil, err
	}

	err = attributevalue.UnmarshalMap(tout.Item, &team)
	if err != nil {
		log.Printf("UnmarshalMapError, %s", err)
		return nil, model.Team{}, nil, err
	}

	tagID := Tags

	var tags []model.Tag = make([]model.Tag, 0)

	for _, v := range tagID {
		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("TAGTABLE")),
			Key: map[string]types.AttributeValue{
				"tag_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Printf("FailedToGetUser, tagID = %v, %s", v, err)
			return nil, model.Team{}, nil, err
		}
		tag := new(model.Tag)
		err = attributevalue.UnmarshalMap(out.Item, &tag)
		if err != nil {
			log.Printf("UnmarshalMapError, %s", err)
			return nil, model.Team{}, nil, err
		}

		tags = append(tags, *tag)
	}

	return users, *team, tags, nil
}
