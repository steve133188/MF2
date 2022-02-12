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

	//find total number of users having this role
	up := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("USERTABLE")),
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("role_id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: id},
		},
	})

	count := 0
	for up.HasMorePages() {
		uout, err := up.NextPage(context.TODO())
		if err != nil {
			fmt.Println("GetRoleByID ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetRoleByID " + err.Error())}), nil
		}
		count += int(uout.Count)
	}

	role := new(model.FullRole)
	role.Total = count
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
	fullRoles := make([]model.FullRole, 0)
	for _, v := range roles {
		// up := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		// 	TableName:        aws.String(os.Getenv("USERTABLE")),
		// 	Limit:            aws.Int32(1000),
		// 	FilterExpression: aws.String("role_id = :id"),
		// 	ExpressionAttributeValues: map[string]types.AttributeValue{
		// 		":id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
		// 	},
		// })

		// count := 0
		// for up.HasMorePages() {
		// 	uout, err := up.NextPage(context.TODO())
		// 	if err != nil {
		// 		fmt.Println("GetAllRoles ", err)
		// 		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetAllRoles " + err.Error())}), nil
		// 	}
		// 	count += int(uout.Count)
		// }

		sout, err := dynaClient.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:              aws.String(os.Getenv("USERTABLE")),
			IndexName:              aws.String("role_id-index"),
			KeyConditionExpression: aws.String("role_id = :rid"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":rid": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
			},
		})
		if err != nil {
			fmt.Println("GetAllRoles ", err)
		}
		// sout, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		// 	TableName:        aws.String(os.Getenv("USERTABLE")),
		// 	FilterExpression: aws.String("role_id = :id"),
		// 	ExpressionAttributeValues: map[string]types.AttributeValue{
		// 		":id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
		// 	},
		// })
		// if err != nil {
		// 	fmt.Println("GetAllRoles ", err)
		// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetAllRoles " + err.Error())}), nil
		// }

		count := int(sout.Count)

		frole := new(model.FullRole)
		val, err := json.Marshal(v)
		if err != nil {
			fmt.Println("GetAllRoles ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetAllRoles " + err.Error())}), nil
		}
		err = json.Unmarshal(val, &frole)
		if err != nil {
			fmt.Println("GetAllRoles ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetAllRoles " + err.Error())}), nil
		}

		frole.Total = count
		fullRoles = append(fullRoles, *frole)
	}

	return ApiResponse(http.StatusOK, fullRoles), nil
}
