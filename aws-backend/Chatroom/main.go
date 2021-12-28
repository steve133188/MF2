package main

import (
	"aws-lambda-chatroom/Handler"
	"aws-lambda-chatroom/Middleware"
	"aws-lambda-chatroom/Services"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	// region     string = "ap-east-1"
	// tableName  string = "MF2_TCO_CHATROOMS"
	dynaClient *dynamodb.Client
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("TABLENAME = ", os.Getenv("TABLENAME"))
	log.Println("REGION = ", os.Getenv("REGION"))

	switch req.HTTPMethod {
	case "GET":
		return Services.GetItems(req, os.Getenv("TABLENAME"), dynaClient)
	case "POST":
		return Services.AddItem(req, os.Getenv("TABLENAME"), dynaClient)
	case "PUT":
		return Services.EditItem(req, os.Getenv("TABLENAME"), dynaClient)
	case "DELETE":
		return Services.DeleteItem(req, os.Getenv("TABLENAME"), dynaClient)
	default:
		return Handler.ApiResponse(http.StatusInternalServerError, map[string]string{
			"message": "InvalidMethod",
		}), nil
	}
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		// o.Region = Utils.GoDotEnvVariable("REGION")
		return nil
	})
	if err != nil {
		log.Println("LoadDefaultConfig    ", err)
		return
	}

	svc := dynamodb.NewFromConfig(cfg)
	dynaClient = svc

	lambda.Start(
		Middleware.Middleware(
			Middleware.JWTHandler(handleRequest),
		),
	)
}
