package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"mf2-aws-dashboard/model"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func UpdateLivechat() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	livechat := new(model.LiveChat)
	livechat.PK = "livechat"
	livechat.TimeStamp = time.Now().Unix()

	timeEnd := livechat.TimeStamp
	timeStart := timeEnd - 24*3600
	fmt.Println(timeStart)

	// user list
	userList := make([]model.User, 0)
	up := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("USERTABLE")),
		Limit:     aws.Int32(100),
	})
	for up.HasMorePages() {
		outs, err := up.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return err
		}

		pUserList := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pUserList)
		if err != nil {
			fmt.Println(err)
			return err
		}

		userList = append(userList, pUserList...)
	}

	// customer list
	customers := make([]model.Customer, 0)
	cp := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("CUSTOMERTABLE")),
		Limit:     aws.Int32(100),
	})

	for cp.HasMorePages() {
		outs, err := cp.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return err
		}

		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pCustomers)
		if err != nil {
			fmt.Println(err)
			return err
		}

		customers = append(customers, pCustomers...)
	}

	// tag list
	tags := make([]model.Tag, 0)
	tp := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TAGTABLE")),
		Limit:     aws.Int32(100),
	})

	for tp.HasMorePages() {
		outs, err := tp.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return err
		}

		pTags := make([]model.Tag, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pTags)
		if err != nil {
			fmt.Println(err)
			return err
		}

		tags = append(tags, pTags...)
	}

	// message list
	messages := make([]model.Message, 0)
	mp := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("MESSAGETABLE")),
		FilterExpression: aws.String("timestamp >= :st AND timestamp <= :et"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":st": &types.AttributeValueMemberN{Value: strconv.FormatInt(timeStart, 10)},
			":et": &types.AttributeValueMemberN{Value: strconv.FormatInt(timeEnd, 10)},
		},
		Limit: aws.Int32(100),
	})

	for mp.HasMorePages() {
		outs, err := mp.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return err
		}

		pMessages := make([]model.Message, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pMessages)
		if err != nil {
			fmt.Println(err)
			return err
		}

		messages = append(messages, pMessages...)
	}

	// sort messages from early time to later time
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].TimeStamp < messages[j].TimeStamp
	})

	// put items into Users
	livechat.Users = make(map[int]model.UserInfo)
	for _, v := range userList {
		userId := v.UserID
		userInfo := new(model.UserInfo)
		userInfo.AllContacts = GetAllContact(userId, customers)
		userInfo.TotalMsgSent = GetTotalMsgSent(userId, messages)
		userInfo.TotalMsgRev = GetTotalMsgRev(userId, messages)
		userInfo.RespTime = GetRespTime(userId, messages)
		userInfo.CommunicationNumber = GetCommunicationNumber(userId, messages)
		userInfo.Tags = GetTags(userId, customers, tags)
		livechat.Users[userId] = *userInfo
	}

	av, err := attributevalue.MarshalMap(&livechat)
	if err != nil {
		log.Printf("FailedToMarshalMap: %s", err)
		return err
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("LIVECHATTABLE")),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(TimeStamp)"),
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemExisted: %s", err)
		}
		log.Printf("ErrorToAddItem: %s", err)
		return err
	}

	return nil
}
