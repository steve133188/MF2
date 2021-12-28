package handler

import (
	"aws-lambda-role/model"
	"aws-lambda-role/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AddRole(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	role := new(model.Role)

	err := json.Unmarshal([]byte(req.Body), &role)
	if err != nil {
		fmt.Println("FailedToUnmarsialReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarsialReqBody, " + err.Error())}), nil
	}

	role.RoleID = utils.IdGenerator()

	av, err := attributevalue.MarshalMap(&role)
	if err != nil {
		fmt.Println("FailedToMarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap, " + err.Error())}), nil
	}

	out, err := dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("TABLE")),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(role_id)"),
		ReturnValues:        types.ReturnValueAllOld,
	})
	if err != nil {
		fmt.Println("FailedToPutItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToPutItem, " + err.Error())}), nil
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &role)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, role), nil
}
