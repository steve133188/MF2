package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func DeleteUserByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	userId := req.PathParameters["id"]
	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: userId},
		},
		ConditionExpression: aws.String("attribute_exists(user_id)"),
	})
	if err != nil {
		fmt.Printf("FailedToDeleteItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func DeleteUsers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	userIDs := req.MultiValueQueryStringParameters["id"]

	for _, v := range userIDs {
		_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: v},
			},
		})
		if err != nil {
			// fmt.Println("FailedToDeleteItem, UserID = ", v, ", ", err)
			fmt.Println(err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem, UserID = " + v + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil
}
