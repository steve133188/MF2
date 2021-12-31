package handler

import (
	"context"
	"fmt"
	"log"
	"mf2-lambda-dashboardv2/model"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func UpdateDashBoard() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	dashboard := new(model.Dashboard)
	dashboard.PK = "PK"
	dashboard.TimeStamp = time.Now().Unix()

	timeEnd := dashboard.TimeStamp
	timeStart := timeEnd - 24*3600
	fmt.Println(timeStart)

	// scan all user list
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

	// scan all customer list
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

	// scan all tag list
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

	// scan all message list
	availableChannel := []string{"whatsapp", "waba"} //current available message channels
	channelMsg := make([]model.ChannelMsg, 0)
	for _, v := range availableChannel {
		messages := make([]model.Message, 0)
		mp := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
			TableName:        aws.String(os.Getenv("MESSAGETABLE")),
			FilterExpression: aws.String("#c = :cg AND #n BETWEEN :st AND :et"),
			ExpressionAttributeNames: map[string]string{
				"#c": "channel",
				"#n": "timestamp",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":cg": &types.AttributeValueMemberS{Value: v},
				":st": &types.AttributeValueMemberS{Value: strconv.FormatInt(timeStart, 10)},
				":et": &types.AttributeValueMemberS{Value: strconv.FormatInt(timeEnd, 10)},
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
		fmt.Println(messages)

		// sort messages from early time to later time
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].TimeStamp < messages[j].TimeStamp
		})

		tempChannelMsg := new(model.ChannelMsg)
		tempChannelMsg.ChannelName = v
		tempChannelMsg.ChannelMessage = messages
		channelMsg = append(channelMsg, *tempChannelMsg)
	}

	// put items into the Dashboard struct
	dashboard.User = make([]model.UserInfo, 0)
	for _, v := range userList {

		// user info
		userInfo := new(model.UserInfo)
		userInfo.UserID = v.UserID
		userInfo.UserName = v.Username
		userInfo.UserRoleName, err = GetRoleName(svc, v.RoleID)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, " RoleID = ", v.RoleID, ", ", err)
			return err
		}
		userInfo.UserStatus = v.Status
		userInfo.LastLogin = v.LastLogin

		//contact info
		count, assignedCustomers := GetAssignedContacts(v.UserID, customers)
		userInfo.AssignedContacts = count

		//channel message info
		userInfo.ChannelData = make([]model.ChannelData, 0)
		for _, k := range channelMsg {
			if len(k.ChannelMessage) == 0 {
				continue
			}

			channelData := new(model.ChannelData)
			channelData.ChannelName = k.ChannelName

			respTime, err := GetRespTime(v.UserID, k.ChannelMessage)
			if err != nil {
				return err
			}
			channelData.AvgRespTime = respTime.Average
			channelData.FirstRespTime = respTime.First
			channelData.LongestRespTime = respTime.Longest

			channelData.MsgSent = GetTotalMsgSent(v.UserID, k.ChannelMessage)
			channelData.MsgRev = GetTotalMsgRev(v.UserID, k.ChannelMessage)
			channelData.CommunicationNumber = GetCommunicationNumber(v.UserID, k.ChannelMessage)

			userInfo.ChannelData = append(userInfo.ChannelData, *channelData)
			//user contact
			userInfo.ActiveContacts += GetActiveContacts(v.UserID, assignedCustomers, k.ChannelMessage)
			//user message info
			userInfo.AvgTotalRespTime += channelData.AvgRespTime
			userInfo.AvgTotalFirstRespTime += channelData.FirstRespTime
			userInfo.TotalMsgSent += channelData.MsgSent
			userInfo.TotalMsgRev += channelData.MsgRev
			userInfo.TotalCommunicationNumber += channelData.CommunicationNumber
		}

		//contact info
		userInfo.UnhandledContact = userInfo.AssignedContacts - userInfo.ActiveContacts
		userInfo.AllContacts = GetAllContact(v.UserID, customers)
		userInfo.Tags = GetTags(v.UserID, customers, tags)

		//user message info
		if len(userInfo.ChannelData) != 0 {
			userInfo.AvgTotalFirstRespTime = userInfo.AvgTotalFirstRespTime / int64(len(userInfo.ChannelData))
			userInfo.AvgTotalRespTime = userInfo.AvgTotalRespTime / int64(len(userInfo.ChannelData))
		}

		dashboard.User = append(dashboard.User, *userInfo)
	}

	fmt.Println(dashboard)

	av, err := attributevalue.MarshalMap(&dashboard)
	if err != nil {
		log.Printf("FailedToMarshalMap: %s", err)
		return err
	}

	// put item to TABLE
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DASHBOARDTABLE")),
		Item:      av,
		//ConditionExpression: aws.String("attribute_not_exists(TimeStamp)"),
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
