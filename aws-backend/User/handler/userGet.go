package handler

import (
	"aws-lambda-user/model"
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

func GetUserByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	userId := req.PathParameters["id"]

	if userId == strconv.Itoa(1) {
		fmt.Println("SystemAccountCannotGet")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("SystemAccountCannotGet")}), nil
	}
	user := new(model.User)

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: userId},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil

	}
	if len(out.Item) == 0 {
		fmt.Println("ItemNotFound")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Item, &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	fullUser := new(model.FullUser)

	err = attributevalue.UnmarshalMap(out.Item, &fullUser)
	if err != nil {
		fmt.Printf("FailedToUnmarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	if user.TeamID != 0 {
		tout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("ORGTABLE")),
			Key: map[string]types.AttributeValue{
				"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.TeamID)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetOrgItem, OrgID = ", user.TeamID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetOrgItem, OrgID = " + strconv.Itoa(user.TeamID) + ", " + err.Error())}), nil
		}
		team := new(model.Team)
		err = attributevalue.UnmarshalMap(tout.Item, &team)
		if err != nil {
			fmt.Println("FailedToUnmarshal, OrgId = , ", user.TeamID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, OrgId = , " + strconv.Itoa(user.TeamID) + ", " + err.Error())}), nil
		}

		fullUser.Team = *team

	}

	if user.RoleID != 0 {
		rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("ROLETABLE")),
			Key: map[string]types.AttributeValue{
				"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(user.RoleID)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetRoleItem, RoleID = ", user.RoleID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + strconv.Itoa(user.RoleID) + ", " + err.Error())}), nil
		}

		role := new(model.Role)
		err = attributevalue.UnmarshalMap(rout.Item, &role)
		if err != nil {
			fmt.Println("FailedToUnmarshal, RoleId = , ", user.RoleID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, RoleId = , " + strconv.Itoa(user.RoleID) + ", " + err.Error())}), nil
		}
		fullUser.RoleName = role.RoleName
		fullUser.Authority = role.Auth
	}

	fullUser.Leads, err = getUserLeads(strconv.Itoa(user.UserID), dynaClient)
	if err != nil {
		fmt.Println("GetUserByID ", err)
		return ApiResponse(http.StatusInternalServerError, aws.String("GetUserByID "+err.Error())), nil
	}

	return ApiResponse(http.StatusOK, fullUser), nil

}

func GetUsersByTeamID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	teamId := req.PathParameters["teamId"]
	Users := make([]model.User, 0)

	p := dynamodb.NewQueryPaginator(dynaClient, &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("team_id-index"),
		KeyConditionExpression: aws.String("team_id = :tid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":tid": &types.AttributeValueMemberN{Value: teamId},
		},
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("FailedToQuery, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToQuery")}), nil
		}
		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pUsers)
		if err != nil {
			fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
		}

		Users = append(Users, pUsers...)
	}

	tout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("ORGTABLE")),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: teamId},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetOrgItem, OrgID = ", teamId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetOrgItem, OrgID = " + teamId + ", " + err.Error())}), nil
	}

	team := new(model.Team)
	err = attributevalue.UnmarshalMap(tout.Item, &team)
	if err != nil {
		fmt.Println("FailedToUnmarshal, OrgId = , ", teamId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, OrgId = , " + teamId + ", " + err.Error())}), nil
	}

	result := make([]model.FullUser, 0)
	for _, v := range Users {
		fUser := new(model.FullUser)
		res, err := json.Marshal(v)
		if err != nil {
			fmt.Println("FailedToMarshalUser, UserID = ", v.UserID)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalUser, UserID = " + strconv.Itoa(v.UserID))}), nil
		}
		err = json.Unmarshal(res, &fUser)
		if err != nil {
			fmt.Println("FailedToUnmarshalUser, UserID = ", v.UserID)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalUser, UserID = " + strconv.Itoa(v.UserID))}), nil
		}

		if v.RoleID != 0 {
			rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
				TableName: aws.String(os.Getenv("ROLETABLE")),
				Key: map[string]types.AttributeValue{
					"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
				},
			})
			if err != nil {
				fmt.Println("FailedToGetOrgItem, RoleID = ", v.RoleID, ", ", err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
			}

			role := new(model.Role)
			err = attributevalue.UnmarshalMap(rout.Item, &role)
			if err != nil {
				fmt.Println("FailedToUnmarshalMap RoleID = ", v.RoleID, ", ", err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
			}
			fUser.RoleName = role.RoleName
			fUser.Authority = role.Auth
		}
		fUser.Leads, err = getUserLeads(strconv.Itoa(v.UserID), dynaClient)
		if err != nil {
			fmt.Println("GetUsersByTeamID ", err)
			return ApiResponse(http.StatusInternalServerError, aws.String("GetUsersByTeamID "+err.Error())), nil
		}

		fUser.Team = *team

		result = append(result, *fUser)
	}

	return ApiResponse(http.StatusOK, result), nil
}

func GetUsersByRoleID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	roleId := req.PathParameters["roleId"]
	var users []model.User = make([]model.User, 0)
	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("role_id = :r AND user_id <> :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":r":  &types.AttributeValueMemberN{Value: roleId},
			":id": &types.AttributeValueMemberN{Value: strconv.Itoa(1)},
		},
	})

	rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("ROLETABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: roleId},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetOrgItem, RoleID = ", roleId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + roleId + ", " + err.Error())}), nil
	}

	role := new(model.Role)
	err = attributevalue.UnmarshalMap(rout.Item, &role)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap RoleID = ", roleId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, RoleID = " + roleId + ", " + err.Error())}), nil
	}

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("FailedToScan, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
		}

		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pUsers)
		if err != nil {
			fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
		}

		users = append(users, pUsers...)
	}

	fullUsers := make([]model.FullUser, 0)
	for _, v := range users {
		fullUser := new(model.FullUser)
		res, err := json.Marshal(v)
		if err != nil {
			fmt.Println("FailedToMarshal, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshal, " + err.Error())}), nil
		}
		err = json.Unmarshal(res, &fullUser)
		if err != nil {
			fmt.Println("FailedToUnmarshal, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, " + err.Error())}), nil
		}
		if v.TeamID != 0 {
			out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
				TableName: aws.String(os.Getenv("ORGTABLE")),
				Key: map[string]types.AttributeValue{
					"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.TeamID)},
				},
			})
			if err != nil {
				fmt.Println("FailedToGetOrg, OrgID = ", v.TeamID, ", ", err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetOrg, OrgID = " + strconv.Itoa(v.TeamID) + ", " + err.Error())}), nil
			}

			attributevalue.UnmarshalMap(out.Item, &fullUser.Team)
		}
		fullUser.RoleName = role.RoleName
		fullUser.Authority = role.Auth
		fullUser.Leads, err = getUserLeads(strconv.Itoa(v.UserID), dynaClient)
		if err != nil {
			fmt.Println("GetUsersByRoleID", err)
			return ApiResponse(http.StatusInternalServerError, aws.String("GetUsersByRoleID"+err.Error())), nil
		}

		fullUsers = append(fullUsers, *fullUser)
	}
	return ApiResponse(http.StatusOK, fullUsers), nil

}

func GetUsers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var users []model.User = make([]model.User, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(1000),
		FilterExpression: aws.String("user_id <> :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: strconv.Itoa(1)},
		},
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
		}

		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pUsers)
		if err != nil {
			fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
		}

		users = append(users, pUsers...)
	}

	// fullUsers := make([]model.FullUser, 0)
	// for _, v := range users {
	// 	fullUser := new(model.FullUser)
	// 	res, err := json.Marshal(v)
	// 	if err != nil {
	// 		fmt.Println("FailedToMarshal, ", err)
	// 		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshal, " + err.Error())}), nil
	// 	}
	// 	err = json.Unmarshal(res, &fullUser)
	// 	if err != nil {
	// 		fmt.Println("FailedToUnmarshal, ", err)
	// 		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, " + err.Error())}), nil
	// 	}
	// 	if v.TeamID != 0 {
	// 		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
	// 			TableName: aws.String(os.Getenv("ORGTABLE")),
	// 			Key: map[string]types.AttributeValue{
	// 				"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.TeamID)},
	// 			},
	// 		})
	// 		if err != nil {
	// 			fmt.Println("FailedToGetOrg, OrgID = ", v.TeamID, ", ", err)
	// 			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetOrg, OrgID = " + strconv.Itoa(v.TeamID) + ", " + err.Error())}), nil
	// 		}
	// 		if out.Item == nil {
	// 			fmt.Println("OrgNotExists, OrgID = ", v.TeamID)
	// 			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("OrgNotExists, OrgID = " + strconv.Itoa(v.TeamID))}), nil
	// 		}
	// 		attributevalue.UnmarshalMap(out.Item, &fullUser.Team)
	// 	}

	// 	if v.RoleID != 0 {
	// 		rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
	// 			TableName: aws.String(os.Getenv("ROLETABLE")),
	// 			Key: map[string]types.AttributeValue{
	// 				"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
	// 			},
	// 		})
	// 		if err != nil {
	// 			fmt.Println("FailedToGetOrgItem, RoleID = ", v.RoleID, ", ", err)
	// 			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
	// 		}
	// 		if rout.Item == nil {
	// 			fmt.Println("OrgNotExists, RoleID = ", v.RoleID)
	// 			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("RoleNotExists, RoleID = " + strconv.Itoa(v.RoleID))}), nil
	// 		}
	// 		role := new(model.Role)
	// 		err = attributevalue.UnmarshalMap(rout.Item, &role)
	// 		if err != nil {
	// 			fmt.Println("FailedToUnmarshalMap RoleID = ", v.RoleID, ", ", err)
	// 			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
	// 		}
	// 		fullUser.RoleName = role.RoleName
	// 		fullUser.Authority = role.Auth
	// 	}
	// 	// fullUser.Leads, err = getUserLeads(strconv.Itoa(v.UserID), dynaClient)
	// 	// if err != nil {
	// 	// 	fmt.Println("GetUsers", err)
	// 	// 	return ApiResponse(http.StatusInternalServerError, aws.String("GetUsers"+err.Error())), nil
	// 	// }
	// 	fullUser.Leads = 0
	// 	fullUsers = append(fullUsers, *fullUser)
	// }

	return ApiResponse(http.StatusOK, users), nil
}

func GetUsersWithoutTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	users := make([]model.User, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("team_id = :val and user_id <> :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
			":id":  &types.AttributeValueMemberN{Value: strconv.Itoa(1)},
		},
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("GetUsersWithoutTeam ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetUsersWithoutTeam " + err.Error())}), nil
		}

		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pUsers)
		if err != nil {
			fmt.Println("GetUsersWithoutTeam ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("GetUsersWithoutTeam " + err.Error())}), nil
		}

		users = append(users, pUsers...)
	}

	fullUsers := make([]model.FullUser, 0)
	for _, v := range users {
		fullUser := new(model.FullUser)
		res, err := json.Marshal(v)
		if err != nil {
			fmt.Println("FailedToMarshal, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshal, " + err.Error())}), nil
		}
		err = json.Unmarshal(res, &fullUser)
		if err != nil {
			fmt.Println("FailedToUnmarshal, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal, " + err.Error())}), nil
		}

		if v.RoleID != 0 {
			rout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
				TableName: aws.String(os.Getenv("ROLETABLE")),
				Key: map[string]types.AttributeValue{
					"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.RoleID)},
				},
			})
			if err != nil {
				fmt.Println("FailedToGetOrgItem, RoleID = ", v.RoleID, ", ", err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetRoleItem, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
			}

			role := new(model.Role)
			err = attributevalue.UnmarshalMap(rout.Item, &role)
			if err != nil {
				fmt.Println("FailedToUnmarshalMap RoleID = ", v.RoleID, ", ", err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, RoleID = " + strconv.Itoa(v.RoleID) + ", " + err.Error())}), nil
			}
			fullUser.RoleName = role.RoleName
			fullUser.Authority = role.Auth
		}
		fullUser.Leads, err = getUserLeads(strconv.Itoa(v.UserID), dynaClient)
		if err != nil {
			fmt.Println("GetUsersWithoutTeam", err)
			return ApiResponse(http.StatusInternalServerError, aws.String("GetUsersWithoutTeam"+err.Error())), nil
		}
		fullUsers = append(fullUsers, *fullUser)
	}

	return ApiResponse(http.StatusOK, fullUsers), nil
}

func GetUserWhatsappInfo(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	userId := req.PathParameters["userId"]

	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        aws.String("whatsapp_nodeTable"),
		FilterExpression: aws.String("user_id = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberN{Value: userId},
		},
	})
	if err != nil {
		fmt.Println("Error in Scan,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("Error in Scan, " + err.Error())}), nil
	}

	node := new(model.Node)
	node.UserId, _ = strconv.Atoi(userId)
	if out.Count == 0 {
		return ApiResponse(http.StatusOK, node), nil
	}
	err = attributevalue.UnmarshalMap(out.Items[0], &node)
	if err != nil {
		fmt.Println("Error in Unmarshal data,", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("Error in Unmarshal data, " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, node), nil
}
