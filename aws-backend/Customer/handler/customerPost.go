package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AddCustomerItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customer := new(model.Customer)

	err := json.Unmarshal([]byte(req.Body), &customer)
	if err != nil {
		log.Printf("FailedToUnmarshalInputData: %s", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalInputData")}), nil
	}

	if customer.CustomerID == 0 {
		customer.CustomerID = customer.Phone
	}
	fmt.Println(customer.CustomerID)
	if customer.CustomerID > 100000000 && customer.CustomerID < 9999999 {
		fmt.Println("UserIdInValid")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserIdInValid")}), nil
	}

	customer.CreatedAt = time.Now().Unix()
	customer.UpdateAt = time.Now().Unix()

	if len(customer.AgentsID) == 0 {
		customer.AgentsID = make([]int, 0)
	} else {
		customer.HandlerId = customer.AgentsID[0]
	}

	if len(customer.Channels) == 0 {
		customer.Channels = make([]string, 0)
	}

	if len(customer.TagsID) == 0 {
		customer.TagsID = make([]int, 0)
	}

	av, err := attributevalue.MarshalMap(&customer)
	if err != nil {
		log.Printf("FailedToMarshalMap: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap")}), nil
	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(customer_id)"),
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemExisted: %s", err)
			return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ItemExisted")}), nil
		}
		log.Printf("ErrorToAddItem: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorToAddItem")}), nil

	}

	// fmt.Println("Adding leads to Agent:", customer.AgentsID)
	// err = ChangeAgentLeads('+', 1, customer.AgentsID, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusCreated, customer), nil
}

func GetCustomerByAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		UserID int `json:"user_id"`
		TeamID int `json:"team_id"`
		RoleID int `json:"role_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("error in unmarshal request body,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in unmarshal request body, " + err.Error())}), nil
	}

	if data.RoleID == 0 || data.UserID == 0 || data.TeamID == 0 {
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("missing input data")}), err
	}

	role := new(model.Role)
	rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("ROLETABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.RoleID)},
		},
	})
	if err != nil {
		fmt.Println("error in get role item,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in get role item, " + err.Error())}), nil
	}
	err = attributevalue.UnmarshalMap(rout.Item, &role)
	if err != nil {
		fmt.Println("error in unmarshal role item,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in unmarshal role item, " + err.Error())}), nil
	}

	var p *dynamodb.ScanPaginator
	if role.Auth.All {
		p = dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
			TableName: aws.String(table),
		})
	} else {
		filterStr := ""
		filterExp := make(map[string]types.AttributeValue)
		if role.Auth.WABA {
			filterStr = "contains(channels, :waba) or "
			filterExp[":waba"] = &types.AttributeValueMemberS{Value: "WABA"}
		}
		if role.Auth.Whatsapp {
			filterStr += "team_id = :tid"
			filterExp[":tid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.TeamID)}
		} else {
			filterStr += "handler_id = :uid"
			filterExp[":uid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)}
		}
		p = dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
			TableName:                 aws.String(table),
			FilterExpression:          aws.String(filterStr),
			ExpressionAttributeValues: filterExp,
		})
		fmt.Println("filterStr = ", filterStr)
	}

	customers := make([]model.Customer, 0)

	count := 0
	for p.HasMorePages() {
		souts, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("error in scanning,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in scanning, " + err.Error())}), nil
		}

		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(souts.Items, &pCustomers)
		if err != nil {
			fmt.Println("error in unmarshal data,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in  unmarshal data, " + err.Error())}), nil
		}
		customers = append(customers, pCustomers...)
		count += int(souts.Count)
	}
	fmt.Println("count = ", count)

	return ApiResponse(http.StatusOK, customers), nil
}
