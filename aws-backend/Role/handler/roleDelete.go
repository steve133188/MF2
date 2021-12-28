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

func DeleteRoleByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		fmt.Println("FailedToDeleteRole, RoleID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteRole, RoleID = " + id + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil

}

func DeleteRoles(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	idList := req.MultiValueQueryStringParameters["id"]

	for _, v := range idList {
		_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(os.Getenv("TABLE")),
			Key: map[string]types.AttributeValue{
				"role_id": &types.AttributeValueMemberN{Value: v},
			},
		})
		if err != nil {
			fmt.Println("FailedToDeleteRole, RoleID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteRole, RoleID = " + v + ", " + err.Error())}), nil
		}
	}
	return ApiResponse(http.StatusOK, nil), nil
}
