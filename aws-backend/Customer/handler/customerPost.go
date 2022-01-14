package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func AddCustomerItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customer := new(model.Customer)

	err := json.Unmarshal([]byte(req.Body), &customer)
	if err != nil {
		log.Printf("FailedToUnmarshalInputData: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalInputData")}), nil
	}

	if customer.CustomerID == 0 {
		customer.CustomerID = customer.Phone
	}
	if customer.CustomerID < 10000000 {
		fmt.Println("UserIdInValid")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserIdInValid")}), nil
	}

	customer.CreatedAt = time.Now().Unix()
	customer.UpdateAt = time.Now().Unix()

	if len(customer.AgentsID) == 0 {
		customer.AgentsID = make([]int, 0)
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
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ItemExisted")}), nil
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

	return ApiResponse(http.StatusOK, customer), nil
}
