package main

import (
	"aws-lambda-chatroom/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"net/http"
	"os"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dynaClient, err := handler.DynamodbConfig()
	if err != nil {
		log.Printf("ErrorToLoadDefaultConfig, %s", err)
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDefaultConfig")}), nil
	}
	log.Println("req.Resource = ", req.Resource)
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/chatroom/{id}":
			return handler.GetChatroomByUser(req, os.Getenv("CHATROOMTABLE"), dynaClient)

		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "DELETE":
		switch req.Resource {
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	default:
		log.Println("InvalidHTTPMethod")
		return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}

}

func main() {
	lambda.Start(handleRequest)
	// 	middleware.Middleware(
	// 		middleware.JWTHandler(handleRequest),
	// 	),
	// )
}
