package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteTagByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"tag_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		fmt.Println("FailedToDeleteTag, TagID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteTag, TagID = " + id + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil

}

func DeleteTags(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	idList := req.MultiValueQueryStringParameters["id"]

	for _, v := range idList {
		_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(os.Getenv("TABLE")),
			Key: map[string]types.AttributeValue{
				"tag_id": &types.AttributeValueMemberN{Value: v},
			},
		})
		if err != nil {
			fmt.Println("FailedToDeleteTag, TagID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteTag, TagID = " + v + ", " + err.Error())}), nil
		}
	}
	return ApiResponse(http.StatusOK, nil), nil
}
