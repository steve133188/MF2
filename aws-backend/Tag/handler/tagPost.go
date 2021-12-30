package handler

import (
	"aws-lambda-tag/model"
	"aws-lambda-tag/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AddTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	tag := new(model.Tag)

	err := json.Unmarshal([]byte(req.Body), &tag)
	if err != nil {
		fmt.Println("FailedToUnmarsialReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarsialReqBody, " + err.Error())}), nil
	}

	tag.TagID = utils.IdGenerator()
	tag.CreatedAt = time.Now().Unix()
	tag.UpdateAt = time.Now().Unix()

	av, err := attributevalue.MarshalMap(&tag)
	if err != nil {
		fmt.Println("FailedToMarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap, " + err.Error())}), nil
	}

	out, err := dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("TABLE")),
		Item:                av,
		ConditionExpression: aws.String("(attribute_not_exists(tag_id)) AND (NOT(contains(#tn, :t)))"),
		ReturnValues:        types.ReturnValueAllOld,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberS{Value: tag.TagName},
		},
		ExpressionAttributeNames: map[string]string{
			"#tn": "tag_name",
		},
	})
	if err != nil {
		fmt.Println("FailedToPutItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToPutItem, " + err.Error())}), nil
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &tag)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, tag), nil
}
