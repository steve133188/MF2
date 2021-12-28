package handler

import (
	"aws-lambda-role/model"
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

func UpdateRoleWithID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	role := new(model.Role)

	err := json.Unmarshal([]byte(req.Body), &role)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody, " + err.Error())}), nil
	}

	updateTime := time.Now().Format(os.Getenv("TIMEFORMAT"))

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(role.RoleID)},
		},
		UpdateExpression: aws.String("Set role_name = :n, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{Value: role.RoleName},
			":t": &types.AttributeValueMemberS{Value: updateTime},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, RoleID = ", role.RoleID, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, RoleID = " + strconv.Itoa(role.RoleID) + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}
