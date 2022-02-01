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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteRoleByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		fmt.Println("FailedToDeleteRole, RoleID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteRole, RoleID = " + id + ", " + err.Error())}), nil
	}

	err = deleteUserRole(dynaClient, id)
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}
	return ApiResponse(http.StatusOK, nil), nil

}

func DeleteRoles(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		RoleID []int `json:"role_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	for _, v := range data.RoleID {
		_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(os.Getenv("TABLE")),
			Key: map[string]types.AttributeValue{
				"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Println("FailedToDeleteRole, RoleID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteRole, RoleID = " + strconv.Itoa(v) + ", " + err.Error())}), nil
		}
		err = deleteUserRole(dynaClient, strconv.Itoa(v))
		if err != nil {
			fmt.Println(err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
		}
	}
	return ApiResponse(http.StatusOK, nil), nil
}

func deleteUserRole(dynaClient *dynamodb.Client, roleId string) error {
	userPage := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("USERTABLE")),
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("role_id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: roleId},
		},
	})

	users := make([]model.User, 0)
	for userPage.HasMorePages() {
		outs, err := userPage.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return err
		}

		pusers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pusers)
		if err != nil {
			fmt.Println(err)
			return err
		}

		users = append(users, pusers...)
	}

	for _, v := range users {
		_, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(os.Getenv("USERTABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.UserID)},
			},
			UpdateExpression: aws.String("Set role_id = :ri"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ri": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateItem, UserID = ", v.UserID, ", ", err)
			return err
		}

	}
	return nil
}
