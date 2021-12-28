package handler

import (
	"aws-lambda-user/model"
	"aws-lambda-user/utils"
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

func UpdateUser(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}
	av, err := attributevalue.MarshalMap(&user)
	if err != nil {
		fmt.Printf("FailedToMarshalMapReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMapReqBody")}), nil
	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           &table,
		Item:                av,
		ConditionExpression: aws.String("attribute_exists(user_id)"),
	})
	if err != nil {
		fmt.Printf("FailedToPutItem, %s", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToPutItem")}), nil
	}

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Item, &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	return ApiResponse(http.StatusOK, user), nil
}

func UpdateUserStatus(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
		},
		UpdateExpression: aws.String("set user_status = :s"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":s": &types.AttributeValueMemberS{Value: user.Status},
		},
		ConditionExpression: aws.String("attribute_exists(user_id)"),
	})
	if err != nil {
		fmt.Printf("FailedToUpdate, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdate")}), nil
	}

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Item, &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	return ApiResponse(http.StatusOK, user), nil
}

func UpdateUserPassword(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		ID          int    `json:"user_id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	user := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.ID)},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	if len(out.Item) == 0 {
		fmt.Printf("UserNotFound, UserID = %v", data.ID)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotFound")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Item, &user)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	match := utils.CheckPasswordHash(data.OldPassword, user.Password)
	if !match {
		fmt.Printf("WrongPassword, UserID = %v", data.ID)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("WrongPassword")}), nil
	}

	user.Password, err = utils.HashPassword(data.NewPassword)
	if err != nil {
		fmt.Println("FailedToHashPassword")
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToHashPassword")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.ID)},
		},
		UpdateExpression: aws.String("set password = :p"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":p": &types.AttributeValueMemberS{Value: user.Password},
		},
	})
	if err != nil {
		fmt.Printf("FailedToUpdateItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func UpdateUserTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
		},
		UpdateExpression: aws.String("set team_id = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberN{Value: strconv.Itoa(user.TeamID)},
		},
		ConditionExpression: aws.String("attribute_exists(user_id)"),
	})
	if err != nil {
		fmt.Printf("FailedToUpdateItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	if user.TeamID != 0 {
		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: &table,
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
			},
		})
		if err != nil {
			fmt.Printf("FailedToGetItem, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
		}
		err = attributevalue.UnmarshalMap(out.Item, &user)
		if err != nil {
			fmt.Printf("FailedToUnmarshalMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
		}
	}

	return ApiResponse(http.StatusOK, user), nil
}

func UpdateUsersTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var updateTeam struct {
		OldTeam int `json:"old_team"`
		NewTeam int `json:"new_team"`
	}

	err := json.Unmarshal([]byte(req.Body), &updateTeam)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody," + err.Error())}), nil
	}
	//checking whether org_id is valid
	if updateTeam.NewTeam != 0 {
		tout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("ORGTABLE")),
			Key: map[string]types.AttributeValue{
				"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(updateTeam.NewTeam)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetOrgItem, OrgID = ", updateTeam.NewTeam, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetOrgItem, OrgID = " + strconv.Itoa(updateTeam.NewTeam) + ", " + err.Error())}), nil
		}
		if tout.Item == nil {
			fmt.Println("OrgNotExists, OrgID = ", updateTeam.NewTeam)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("OrgNotExists, OrgID = " + strconv.Itoa(updateTeam.NewTeam))}), nil
		}
	}

	tp := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("team_id = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberN{Value: strconv.Itoa(updateTeam.OldTeam)},
		},
	})
	users := make([]model.User, 0)
	for tp.HasMorePages() {
		touts, err := tp.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("FailedToScan, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan, " + err.Error())}), nil
		}
		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(touts.Items, &pUsers)
		if err != nil {
			fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
		}

		users = append(users, pUsers...)
	}

	for _, v := range users {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(os.Getenv("TABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.UserID)},
			},
			UpdateExpression: aws.String("Set team_id = :nt"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":nt": &types.AttributeValueMemberN{Value: strconv.Itoa(updateTeam.NewTeam)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateItem, UserID = ", v.UserID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, UserID = " + strconv.Itoa(v.UserID) + ", " + err.Error())}), nil
		}
	}
	return ApiResponse(http.StatusOK, nil), nil
}

func UpdateUserRole(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.UserID)},
		},
		UpdateExpression: aws.String("Set role_id = :ri"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ri": &types.AttributeValueMemberN{Value: strconv.Itoa(user.RoleID)},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, UserID = ", user.UserID, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, UserID = " + strconv.Itoa(user.UserID) + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, user), nil
}

func UpdateUsersRole(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var updateRole struct {
		OldRole int `json:"old_role"`
		NewRole int `json:"new_role"`
	}

	err := json.Unmarshal([]byte(req.Body), &updateRole)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody," + err.Error())}), nil
	}
	//checking whether role_id is valid
	if updateRole.NewRole != 0 {
		tout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("ROLETABLE")),
			Key: map[string]types.AttributeValue{
				"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(updateRole.NewRole)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetRoleItem, RoleID = ", updateRole.NewRole, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + strconv.Itoa(updateRole.NewRole) + ", " + err.Error())}), nil
		}
		if tout.Item == nil {
			fmt.Println("RoleNotExists, RoleID = ", updateRole.NewRole)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("RoleNotExists, RoleID = " + strconv.Itoa(updateRole.NewRole))}), nil
		}
	}

	tp := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("role_id = :r"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":r": &types.AttributeValueMemberN{Value: strconv.Itoa(updateRole.OldRole)},
		},
	})
	users := make([]model.User, 0)
	for tp.HasMorePages() {
		touts, err := tp.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("FailedToScan, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan, " + err.Error())}), nil
		}
		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(touts.Items, &pUsers)
		if err != nil {
			fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
		}

		users = append(users, pUsers...)
	}

	for _, v := range users {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(os.Getenv("TABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.UserID)},
			},
			UpdateExpression: aws.String("Set role_id = :nr"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":nr": &types.AttributeValueMemberN{Value: strconv.Itoa(updateRole.NewRole)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateItem, UserID = ", v.UserID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, UserID = " + strconv.Itoa(v.UserID) + ", " + err.Error())}), nil
		}
	}
	return ApiResponse(http.StatusOK, nil), nil
}

func AddUserChannels(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		UserID   int        `json:"user_id"`
		Channels model.Chan `json:"channels"`
	}
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("UserID = ", data.UserID, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UserID = " + strconv.Itoa(data.UserID) + err.Error())}), nil
	}

	gout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}
	if gout.Item == nil {
		fmt.Println("UserNotExist, UserID = ", data.UserID)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist, UserID = " + strconv.Itoa(data.UserID))}), nil
	}

	user := new(model.User)
	err = attributevalue.UnmarshalMap(gout.Item, &user)
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	//checking if channels is existed
	for _, v := range user.Channels {
		if v.ChannelName == data.Channels.ChannelName {
			fmt.Println("ChannelExist, ChannelName = ", data.Channels.ChannelName)
			return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ChannelExist, ChannelName = " + data.Channels.ChannelName)}), nil
		}
	}

	user.Channels = append(user.Channels, data.Channels)

	lav, err := attributevalue.MarshalList(user.Channels)
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	fmt.Println("lav = ", lav)
	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
		UpdateExpression: aws.String("Set channels = :ch"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ch": &types.AttributeValueMemberL{Value: lav},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, out.Attributes), nil
}

func EditUserChannels(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		UserID   int        `json:"user_id"`
		Channels model.Chan `json:"channels"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("UserID = ", data.UserID, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UserID = " + strconv.Itoa(data.UserID) + err.Error())}), nil
	}

	gout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}
	if gout.Item == nil {
		fmt.Println("UserNotExist, UserID = ", data.UserID)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist, UserID = " + strconv.Itoa(data.UserID))}), nil
	}

	user := new(model.User)
	err = attributevalue.UnmarshalMap(gout.Item, &user)
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	found := false
	var key int
	for k, v := range user.Channels {
		if v.ChannelName == data.Channels.ChannelName {
			found = true
			key = k
		}
	}

	if !found {
		fmt.Println("ChannelNotExist, ChannelName = ", data.Channels.ChannelName)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ChannelNotExist, ChannelName = " + data.Channels.ChannelName)}), nil
	}

	updateStr := "Set channels[" + strconv.Itoa(key) + "].channel_url = :url"

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
		UpdateExpression: aws.String(updateStr),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":url": &types.AttributeValueMemberS{Value: data.Channels.ChannelUrl},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, out.Attributes), nil

}

func DeleteUserChannels(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		UserID   int        `json:"user_id"`
		Channels model.Chan `json:"channels"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("UserID = ", data.UserID, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UserID = " + strconv.Itoa(data.UserID) + err.Error())}), nil
	}

	gout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}
	if gout.Item == nil {
		fmt.Println("UserNotExist, UserID = ", data.UserID)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist, UserID = " + strconv.Itoa(data.UserID))}), nil
	}

	user := new(model.User)
	err = attributevalue.UnmarshalMap(gout.Item, &user)
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	found := false
	var key int
	for k, v := range user.Channels {
		if v.ChannelName == data.Channels.ChannelName {
			found = true
			key = k
		}
	}
	if !found {
		fmt.Println("ChannelNotExist, ChannelName = ", data.Channels.ChannelName)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ChannelNotExist, ChannelName = " + data.Channels.ChannelName)}), nil
	}

	updateStr := "REMOVE channels[" + strconv.Itoa(key) + "]"

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)},
		},
		UpdateExpression: aws.String(updateStr),

		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		fmt.Println(err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, out.Attributes), nil
}
