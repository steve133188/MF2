package handler

import (
	"aws-lambda-user-without-auth/model"
	"aws-lambda-user-without-auth/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dgrijalva/jwt-go"
)

var token_pwd string = "51c3d3fc-3e15-4c19-7437-d74f5e5f906c"

func AddUser(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	if len(user.Email) == 0 && len(user.Password) == 0 {
		fmt.Println("MissingPasswordOrEmail")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingPasswordOrEmail")}), nil
	}

	if user.UserID == 0 {
		user.UserID = utils.IdGenerator()
	}

	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("FailedToHasPassword")
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToHasPassword")}), nil
	}

	av, err := attributevalue.MarshalMap(&user)
	if err != nil {
		fmt.Printf("FailedToMarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap")}), nil
	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(user_id) AND attribute_not_exists(email) AND attribute_not_exists(username)"),
	})
	if err != nil {
		fmt.Printf("FailedToAddItem, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToAddItem")}), nil
	}

	return ApiResponse(http.StatusOK, user), nil
}

func UserLogin(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	user := new(model.User)
	find := new(model.User)

	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		fmt.Println("Login: Unmarshal    ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalData")}), nil
	}
	if len(user.Email) == 0 || len(user.Password) == 0 {
		fmt.Println("Login: Missing Email or Password")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingEmailOrPassword")}), nil
	}

	// out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
	// 	TableName: &tableName,
	// 	Key: map[string]types.AttributeValue{
	// 		"email": &types.AttributeValueMemberS{Value: string(user.Email)},
	// 	},
	// })

	out, err := dynaClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("email-index"),
		KeyConditionExpression: aws.String("email = :emailVal"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":emailVal": &types.AttributeValueMemberS{Value: user.Email},
		},
	})

	if err != nil {
		fmt.Println("Login: FailedToGetUser    ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetUser")}), nil
	}

	if len(out.Items) == 0 {
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Items[0], &find)
	if err != nil {
		fmt.Println("Login: Error in UnmarshalMap    ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalData")}), nil
	}

	passwordCheck := utils.CheckPasswordHash(user.Password, find.Password)
	if !passwordCheck {
		fmt.Println("Login: Wrong Password    ", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("WrongPassword")}), nil
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["username"] = find.Username
	claims["user_id"] = find.UserID
	claims["email"] = find.Email
	claims["role_id"] = find.RoleID
	claims["status"] = find.Status

	s, err := token.SignedString([]byte(token_pwd))
	if err != nil {
		fmt.Println("Login: FailedToSignedToken    ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToSignedToken")}), nil
	}

	return ApiResponse(http.StatusOK, map[string]interface{}{
		"token": s,
		"user":  &find,
	}), nil

}
func UserForgotPassword(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var address struct {
		Address string `json:"address"`
	}
	user := new(model.User)
	err := json.Unmarshal([]byte(req.Body), &address)
	if err != nil {
		fmt.Printf("FailedToUnmarshal, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshal")}), nil
	}

	out, err := dynaClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("email-index"),
		KeyConditionExpression: aws.String("email = :emailVal"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":emailVal": &types.AttributeValueMemberS{Value: address.Address},
		},
	})
	if err != nil {
		fmt.Printf("FailedToGetUser, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetUser")}), nil
	}
	if len(out.Items) == 0 {
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist")}), nil
	}

	err = attributevalue.UnmarshalMap(out.Items[0], &user)
	if err != nil {
		fmt.Printf("FailedToUnmarshalMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}
	fmt.Println("user = ", user)
	randomPassword := utils.GeneratePassword(2, 2, 2, 8)
	password, err := utils.HashPassword(randomPassword)
	if err != nil {
		fmt.Printf("FailedToHasPassword, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToHasPassword")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberN{Value: string(user.UserID)},
		},
		UpdateExpression: aws.String("set password = :pw"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pw":    &types.AttributeValueMemberS{Value: password},
			":email": &types.AttributeValueMemberS{Value: user.Email},
		},
		ConditionExpression: aws.String("email = :email"),
	})
	if err != nil {
		fmt.Printf("FailedToUpdatePassword, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdatePassword")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

// func AddUsers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
// 	var users []model.User = make([]model.User, 0)

// 	err := json.Unmarshal([]byte(req.Body), &users)
// 	if err != nil {
// 		fmt.Printf("FailedToUnmarshalReqBody, %s")
// 		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
// 	}

// 	out, err := dynaClient.PutItem(put)
// }
