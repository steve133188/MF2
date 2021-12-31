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
	dashboard.Channel = make([]model.Channel, 0)
	for _, v := range channelMsg {
		if len(v.ChannelMessage) == 0 {
			continue
		}

		channelData := new(model.Channel)
		channelData.ChannelName = v.ChannelName

		channelData.User = make([]model.UserInfo, 0)
		for _, k := range userList {

			// user info
			userInfo := new(model.UserInfo)
			userInfo.UserID = k.UserID
			userInfo.UserName = k.Username
			userInfo.UserRoleName, err = GetRoleName(svc, k.RoleID)
			if err != nil {
				fmt.Println("UserID = ", k.UserID, " RoleID = ", k.RoleID, ", ", err)
				return err
			}
			userInfo.UserStatus = k.Status
			userInfo.LastLogin = k.LastLogin

			//contact info
			count, assignedCustomers := GetAssignedContacts(k.UserID, customers)
			userInfo.AssignedContacts = count
			userInfo.ActiveContacts += GetActiveContacts(k.UserID, assignedCustomers, v.ChannelMessage)
			userInfo.UnhandledContact = userInfo.AssignedContacts - userInfo.ActiveContacts
			userInfo.AllContacts = GetAllContact(k.UserID, customers)
			userInfo.Tags = GetTags(k.UserID, customers, tags)

			//user dashboard info
			respTime, err := GetRespTime(k.UserID, v.ChannelMessage)
			if err != nil {
				return err
			}
			userInfo.AvgRespTime = respTime.Average
			userInfo.FirstRespTime = respTime.First
			userInfo.LongestRespTime = respTime.Longest

			userInfo.MsgSent = GetTotalMsgSent(k.UserID, v.ChannelMessage)
			userInfo.MsgRev = GetTotalMsgRev(k.UserID, v.ChannelMessage)
			userInfo.CommunicationNumber = GetCommunicationNumber(k.UserID, v.ChannelMessage)

			//channel dashboard info
			channelData.AvgTotalRespTime += userInfo.AvgRespTime
			channelData.AvgTotalFirstRespTime += userInfo.FirstRespTime
			channelData.TotalMsgSent += userInfo.MsgSent
			channelData.TotalMsgRev += userInfo.MsgRev
			channelData.TotalCommunicationNumber += userInfo.CommunicationNumber

			channelData.User = append(channelData.User, *userInfo)
		}

		channelData.AvgTotalFirstRespTime = channelData.AvgTotalFirstRespTime / int64(len(channelData.User))
		channelData.AvgTotalRespTime = channelData.AvgTotalRespTime / int64(len(channelData.User))

		dashboard.Channel = append(dashboard.Channel, *channelData)
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
