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
	fmt.Println("Successful Scan User")

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
	fmt.Println("Successful Scan Customer")

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
	fmt.Println("Successful Scan Tag")

	// scan all message list
	availableChannel := []string{"WhatsApp", "WABA"} //current available message channels
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

		// sort messages from early time to later time
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].TimeStamp < messages[j].TimeStamp
		})

		tempChannelMsg := new(model.ChannelMsg)
		tempChannelMsg.ChannelName = v
		tempChannelMsg.ChannelMessage = messages
		channelMsg = append(channelMsg, *tempChannelMsg)
	}
	fmt.Println(channelMsg)
	fmt.Println("Successful Scan Message")

	// put items into the Dashboard struct
	dashboard.Channel = make([]model.Channel, 0)
	dashboard.User = make([]model.UserInfo, 0)
	for i, v := range channelMsg {
		if len(v.ChannelMessage) == 0 {
			continue
		}

		channelData := new(model.Channel)
		channelData.ChannelName = v.ChannelName

		channelData.User = make([]model.UserInfo, 0)
		for j, k := range userList {

			// user info
			userInfo := new(model.UserInfo)
			userInfo.UserID = k.UserID
			userInfo.UserName = k.Username
			userInfo.UserRoleName, err = GetRoleName(svc, k.RoleID)
			if err != nil {
				fmt.Println("UserID = ", k.UserID, " RoleID = ", k.RoleID, ", ", err)
				return err
			}
			userInfo.TeamID = k.TeamID
			userInfo.UserStatus = k.Status
			userInfo.LastLogin = k.LastLogin

			//user dashboard info
			respTime, activeCount, err := GetRespTime(k.UserID, v.ChannelMessage)
			if err != nil {
				return err
			}
			userInfo.AvgRespTime = respTime.Average
			userInfo.FirstRespTime = respTime.First
			userInfo.LongestRespTime = respTime.Longest

			userInfo.MsgSent = GetMsgSent(k.UserID, v.ChannelMessage)
			userInfo.MsgRev = GetMsgRev(k.UserID, v.ChannelMessage)
			//userInfo.CommunicationNumber = GetCommunicationNumber(k.UserID, v.ChannelMessage)

			//contact info

			count, assignedCustomers := GetAssignedContacts(k.UserID, customers)
			userInfo.AssignedContacts = count
			userInfo.ActiveContacts = activeCount
			userInfo.DeliveredContacts = GetDeliveredContacts(k.UserID, assignedCustomers, v.ChannelMessage)
			userInfo.UnhandledContact = userInfo.AssignedContacts - userInfo.ActiveContacts - userInfo.DeliveredContacts

			userInfo.NewAddedContacts = 0 //later use
			userInfo.AllContacts = userInfo.AssignedContacts + userInfo.NewAddedContacts
			userInfo.Tags = GetTags(k.UserID, customers, tags)

			//channel dashboard info
			channelData.AvgTotalRespTime += userInfo.AvgRespTime
			channelData.AvgTotalFirstRespTime += userInfo.FirstRespTime
			channelData.AvgLongestRespTime += userInfo.LongestRespTime
			//channelData.TotalCommunicationNumber += userInfo.CommunicationNumber

			channelData.NewAddedContacts += userInfo.NewAddedContacts

			channelData.TotalAssignedContacts += userInfo.AssignedContacts
			channelData.TotalActiveContacts += userInfo.ActiveContacts
			channelData.TotalDeliveredContacts += userInfo.DeliveredContacts
			channelData.TotalUnhandledContact += userInfo.UnhandledContact

			channelData.User = append(channelData.User, *userInfo)

			//put item into dashboard user struct
			if i == 0 || len(channelMsg[0].ChannelMessage) == 0 {
				dashboard.User = append(dashboard.User, *userInfo)
			} else {
				dashboard.User[j].AvgRespTime += userInfo.AvgRespTime
				dashboard.User[j].FirstRespTime += userInfo.FirstRespTime
				dashboard.User[j].LongestRespTime += userInfo.LongestRespTime

				dashboard.User[j].MsgSent += userInfo.MsgSent
				dashboard.User[j].MsgRev += userInfo.MsgRev

				dashboard.User[j].AssignedContacts += userInfo.AssignedContacts
				dashboard.User[j].ActiveContacts += userInfo.ActiveContacts
				dashboard.User[j].DeliveredContacts += userInfo.DeliveredContacts
				dashboard.User[j].UnhandledContact += userInfo.UnhandledContact

				dashboard.User[j].NewAddedContacts += userInfo.NewAddedContacts
				dashboard.User[j].AllContacts += userInfo.AllContacts
			}
		}

		channelData.AllContacts = channelData.TotalAssignedContacts + channelData.NewAddedContacts
		channelData.CommunicationHours = GetCommunicationHour(timeStart, timeEnd, 12, v.ChannelMessage)

		channelData.TotalMsgSent = GetTotalMsgSent(v.ChannelMessage)
		channelData.TotalMsgRev = len(v.ChannelMessage) - channelData.TotalMsgSent

		// average the total time
		channelData.AvgTotalFirstRespTime = channelData.AvgTotalFirstRespTime / int64(len(channelData.User))
		channelData.AvgTotalRespTime = channelData.AvgTotalRespTime / int64(len(channelData.User))
		channelData.AvgLongestRespTime = channelData.AvgLongestRespTime / int64(len(channelData.User))

		dashboard.Channel = append(dashboard.Channel, *channelData)

		//Dashboard info
		dashboard.AvgTotalRespTime += channelData.AvgTotalFirstRespTime
		dashboard.AvgTotalFirstRespTime += channelData.AvgTotalFirstRespTime
		dashboard.AvgLongestRespTime += channelData.AvgLongestRespTime

		dashboard.TotalMsgSent += channelData.TotalMsgSent
		dashboard.TotalMsgRev += channelData.TotalMsgRev

		dashboard.CommunicationHours = make([]int, 12)
		for i, c := range channelData.CommunicationHours {
			dashboard.CommunicationHours[i] += c
		}

		dashboard.AllContacts += channelData.AllContacts
		dashboard.NewAddedContacts += channelData.NewAddedContacts
		dashboard.TotalAssignedContacts += channelData.TotalAssignedContacts
		dashboard.TotalActiveContacts += channelData.TotalActiveContacts
		dashboard.TotalDeliveredContacts += channelData.TotalDeliveredContacts
		dashboard.TotalUnhandledContact += channelData.TotalUnhandledContact

		fmt.Println("Done for Channel:", channelData.ChannelName)
	}

	dashboard.AvgTotalFirstRespTime = dashboard.AvgTotalFirstRespTime / int64(len(dashboard.Channel))
	dashboard.AvgTotalRespTime = dashboard.AvgTotalRespTime / int64(len(dashboard.Channel))
	dashboard.AvgLongestRespTime = dashboard.AvgLongestRespTime / int64(len(dashboard.Channel))

	dashboard.Tags = GetAllTags(userList, customers, tags)

	fmt.Println("Done for Dashboard")
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
