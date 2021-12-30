package handler

import (
	"aws-lambda-tag/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	tag := new(model.Tag)

	err := json.Unmarshal([]byte(req.Body), &tag)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody, " + err.Error())}), nil
	}

	updateTime := time.Now().Unix()

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"tag_id": &types.AttributeValueMemberN{Value: strconv.Itoa(tag.TagID)},
		},
		UpdateExpression: aws.String("Set tag_name = :n, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{Value: tag.TagName},
			":t": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, TagID = ", tag.TagID, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, TagID = " + strconv.Itoa(tag.TagID) + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}
