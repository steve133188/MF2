package main

import (
	"aws-lambda-user-without-auth/handler"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	cstZone := time.FixedZone("CST", 8*3600)
	time.Local = cstZone
	dynaClient, err := handler.DynamodbConfig()
	if err != nil {
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDynamodbConfig")}), nil
	}

	table := os.Getenv("TABLE")

	switch req.HTTPMethod {
	case "POST":
		switch req.Resource {
		case "/api/users/login":
			return handler.UserLogin(req, table, dynaClient)
		case "/api/users":
			return handler.AddUser(req, table, dynaClient)
		case "/api/users/forgot-password":
			return handler.UserForgotPassword(req, table, dynaClient)
			// case "/users":
			// 	return handler.AddUsers(req, table, dynaClient)
			// }
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidPostMethod")}), nil
		}
	default:
		return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil

	}
}

func main() {
	lambda.Start(handleRequest)
}
