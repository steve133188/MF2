package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		customer.CustomerID, err = strconv.Atoi(customer.Phone)
		if err != nil {
			log.Printf("WrongIDFormat: %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("WrongIDFormat")}), nil
		}
	}
	customer.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	customer.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

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

	return ApiResponse(http.StatusOK, customer), nil
}
