package handler

import (
	"context"
	"fmt"
	"log"
	"mf2-aws-dashboard/model"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateAgent() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	agent := new(model.Agent)
	agent.PK = "agent"
	agent.TimeStamp = time.Now().Unix()

	timeEnd := agent.TimeStamp
	timeStart := timeEnd - 24*3600
	fmt.Println("timeEnd = ", timeEnd)
	fmt.Println("timeStart = ", timeStart)

	//user list
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

	for _, v := range userList {
		agentInfo := new(model.AgentInfo)
		agentInfo.UserName = v.Username
		agentInfo.UserRoleName, err = GetRoleName(svc, v.RoleID)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, " RoleID = ", v.RoleID, ", ", err)
			return err
		}
		agentInfo.UserStatus = v.Status
		agentInfo.LastLogin = v.LastLogin
		agentInfo.AssignedContacts, err = GetAssignedContacts(svc, v.UserID)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, ", ", err)
			return err
		}
		agentInfo.ActiveContacts, err = GetActiveContacts(svc, v.UserID, timeStart, timeEnd)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, ", ", err)
			return err
		}
		agentInfo.UnhandledContact, err = GetUnhandlerContact(svc, v.UserID, timeStart, timeEnd)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, ", ", err)
			return err
		}
		agentInfo.TotalMsgSent, err = TotalMsgSent(svc, v.UserID, timeStart, timeEnd)
		if err != nil {
			fmt.Println("UserID = ", v.UserID, ", ", err)
			return err
		}
	}
}

func GetRoleName(dynaClient *dynamodb.Client, roleId int) (string, error) {
	role := new(model.Role)
	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("ROLETABLE"),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(roleId)},
		},
	})
	if err != nil {
		return "", err
	}

	err = attributevalue.UnmarshalMap(out.Item, &role)
	if err != nil {
		return "", err
	}

	return role.RoleName, nil
}

func GetAssignedContacts(dynaClient *dynamodb.Client, userId int) (int, error) {
	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String("CUSTOMERTABLE"),
		Limit:            aws.Int32(100),
		FilterExpression: aws.String("contains(agent_id, :userId)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberN{Value: strconv.Itoa(userId)},
		},
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			return 0, err
		}

		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pCustomers)
		if err != nil {
			return 0, err
		}

		customers = append(customers, pCustomers...)
	}

	return len(customers), nil
}

func GetActiveContacts(dynaClient *dynamodb.Client, userId int, start int64, end int64) (int, error) {
	assignedCustomer := make([]model.Customer, 0)
	cp := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String("CUSTOMERTABLE"),
		FilterExpression: aws.String("contains(agents_id, :userId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberN{Value: strconv.Itoa(userId)},
		},
	})
	for cp.HasMorePages() {
		couts, err := cp.NextPage(context.TODO())
		if err != nil {
			return 0, nil
		}
		pAssignedCustomer := make([]model.Customer, 0)

		err = attributevalue.UnmarshalListOfMaps(couts.Items, &pAssignedCustomer)
		if err != nil {
			return 0, nil
		}
		assignedCustomer = append(assignedCustomer, pAssignedCustomer...)
	}

	count := 0
	for _, v := range assignedCustomer {
		out, err := dynaClient.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:              aws.String("MESSAGETABLE"),
			KeyConditionExpression: aws.String("room_id = :PK, #SK BETWEEN :ts AND :te"),
			FilterExpression:       aws.String("receiver = :customerId"),
			ExpressionAttributeNames: map[string]string{
				"#SK": "timestamp",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":PK":         &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
				":ts":         &types.AttributeValueMemberN{Value: strconv.FormatInt(start, 10)},
				":te":         &types.AttributeValueMemberN{Value: strconv.FormatInt(end, 10)},
				":customerId": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
		})
		if err != nil {
			return 0, nil
		}
		if len(out.Items) != 0 {
			count++
		}
	}
	return count, nil
}

func GetUnhandlerContact(dynaClient *dynamodb.Client, userId int, start int64, end int64) (int, error) {
	assignedCustomer := make([]model.Customer, 0)
	cp := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String("CUSTOMERTABLE"),
		FilterExpression: aws.String("contains(agents_id, :userId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberN{Value: strconv.Itoa(userId)},
		},
	})
	for cp.HasMorePages() {
		couts, err := cp.NextPage(context.TODO())
		if err != nil {
			return 0, nil
		}
		pAssignedCustomer := make([]model.Customer, 0)

		err = attributevalue.UnmarshalListOfMaps(couts.Items, &pAssignedCustomer)
		if err != nil {
			return 0, nil
		}
		assignedCustomer = append(assignedCustomer, pAssignedCustomer...)
	}
	count := 0
	for _, v := range assignedCustomer {
		out, err := dynaClient.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:              aws.String("MESSAGETABLE"),
			KeyConditionExpression: aws.String("room_id = :PK, #SK BETWEEN :ts AND :te"),
			FilterExpression:       aws.String("receiver = :customerId"),
			ExpressionAttributeNames: map[string]string{
				"#SK": "timestamp",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":PK":         &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
				":ts":         &types.AttributeValueMemberN{Value: strconv.FormatInt(start, 10)},
				":te":         &types.AttributeValueMemberN{Value: strconv.FormatInt(end, 10)},
				":customerId": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
		})
		if err != nil {
			return 0, nil
		}
		if len(out.Items) != 0 {
			count++
		}
	}
	return len(assignedCustomer) - count, nil
}

func TotalMsgSent(dynaClient *dynamodb.Client, userId int, start int64, end int64) (int, error) {
	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String("MESSAGETABLE"),
		FilterExpression: aws.String("sender = :sen, from_me = true"),
		Limit:            aws.Int32(100),
	})

	msg := make([]model.Message, 0)
	for p.HasMorePages() {
		outs, err := p.NextPage(context.TODO())
		if err != nil {
			return 0, nil
		}
		pmsg := make([]model.Message, 0)

		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pmsg)
		if err != nil {
			return 0, nil
		}
		msg = append(msg, pmsg...)
	}

	return len(msg), nil
}
