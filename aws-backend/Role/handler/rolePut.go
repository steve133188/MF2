package handler

import (
	"aws-lambda-role/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func UpdateRoleWithID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	role := new(model.Role)

	err := json.Unmarshal([]byte(req.Body), &role)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody, " + err.Error())}), nil
	}

	av, err := attributevalue.MarshalMap(role)
	if err != nil {
		fmt.Println("UpdateRoleWithID ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UpdateRoleWithID " + err.Error())}), nil
	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("TABLE")),
		Item:                av,
		ConditionExpression: aws.String("attribute_exists(role_id)"),
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, RoleID = ", role.RoleID, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, RoleID = " + strconv.Itoa(role.RoleID) + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}
