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

func AddStdReply(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	stdReply := new(model.StandardReplies)

	err := json.Unmarshal([]byte(req.Body), &stdReply)
	if err != nil {
		fmt.Println("FailedToUnmarsialReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarsialReqBody, " + err.Error())}), nil
	}

	if stdReply.ID == "" {
		id := time.Now().UnixNano() / 1e6
		stdReply.ID = strconv.FormatInt(id, 10)
	}

	if len(stdReply.Body) == 0 {
		stdReply.Body = make([]string, 0)
	}

	if len(stdReply.Variables) == 0 {
		stdReply.Variables = make([]string, 0)
	}

	if len(stdReply.Channels) == 0 {
		stdReply.Channels = make([]string, 0)
	}

	stdReply.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	stdReply.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	av, err := attributevalue.MarshalMap(&stdReply)
	if err != nil {
		fmt.Println("FailedToMarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap, " + err.Error())}), nil
	}

	out, err := dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
		ReturnValues:        types.ReturnValueAllOld,
	})
	if err != nil {
		fmt.Println("FailedToPutItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToPutItem, " + err.Error())}), nil
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &stdReply)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, stdReply), nil
}
