package handler

import (
	"context"
	"fmt"
	"log"
	"mf2-aws-dashboard/model"
	"os"
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

	// message list
	messages := make([]model.Message, 0)
	mp := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("MESSAGETABLE")),
		Limit:     aws.Int32(100),
	})

	for mp.HasMorePages() {
		outs, err := cp.NextPage(context.TODO())
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

	// put items into Users

	for _, v := range userList {
		userId := v.UserID
		userInfo := new(model.UserInfo)
		userInfo.AllContacts = GetAllContact(userId, customers)

	}

	livechat.Users = make(map[int]model.UserInfo)

}

func GetAllContact(userID int, customers []model.Customer) int {
	var count int
	for _, v := range customers {
		for _, vAgent := range v.AgentsID {
			if vAgent == userID {
				count++
			}
		}
	}
	return count
}
