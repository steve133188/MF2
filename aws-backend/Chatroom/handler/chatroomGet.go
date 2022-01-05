package handler

import (
	"aws-lambda-chatroom/model"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
)

func GetChatroomByUser(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	userId := req.PathParameters["id"]

	chatrooms := make([]model.Chatroom, 0)

	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("user_id = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberN{Value: userId},
		},
	})

	if err != nil {
		fmt.Printf("FailedToScan, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &chatrooms)
	if err != nil {
		fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
	}

	return ApiResponse(http.StatusOK, chatrooms), nil
}
