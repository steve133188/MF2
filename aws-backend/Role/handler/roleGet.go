package handler

import (
	"aws-lambda-role/model"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetRoleByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, RoleID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem, RoleID = " + id + ", " + err.Error())}), nil
	}
	if out.Item == nil {
		fmt.Println("RoleNotExists, RoleID = ", id, ", ", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("RoleNotExists, RoleID = " + id + ", " + err.Error())}), nil
	}

	role := new(model.Role)
	err = attributevalue.UnmarshalMap(out.Item, &role)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, RoleID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, RoleID = " + id + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, role), nil
}

func GetAllRoles(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	roles := make([]model.Role, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Limit:     aws.Int32(50),
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan, " + err.Error())}), nil
		}

		pRoles := make([]model.Role, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pRoles)
		if err != nil {
			fmt.Println("FailedToUnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMaps, " + err.Error())}), nil
		}

		roles = append(roles, pRoles...)
	}

	return ApiResponse(http.StatusOK, roles), nil
}
