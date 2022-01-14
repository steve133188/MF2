package handler

import (
	"aws-lambda-standard-replies/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateStdReplyByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	stdReply := new(model.StandardReplies)

	err := json.Unmarshal([]byte(req.Body), &stdReply)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody, " + err.Error())}), nil
	}

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: stdReply.ID},
		},
	})
	if err != nil {
		fmt.Println("failed to get item, ID = ", stdReply.ID)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to get item, ID = " + stdReply.ID)}), nil
	}
	if out.Item == nil {
		fmt.Println("item not existed, ID = ", stdReply.ID)
		return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("item not existed, ID = " + stdReply.ID)}), nil
	}

	data := new(model.StandardReplies)
	err = attributevalue.UnmarshalMap(out.Item, &data)
	if err != nil {
		fmt.Println("failed to UnmarshalMap item, ID = ", stdReply.ID)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to UnmarshalMap item, ID = " + stdReply.ID)}), nil
	}

	stdReply.CreatedAt = data.CreatedAt
	stdReply.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	av, err := attributevalue.MarshalMap(stdReply)
	if err != nil {
		fmt.Println("failed to MarshalMap item, ID = ", stdReply.ID)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to MarshalMap item, ID = " + stdReply.ID)}), nil
	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      av,
	})
	if err != nil {
		fmt.Println("failed to add item, ID = ", stdReply.ID)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to add item, ID = " + stdReply.ID)}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}
