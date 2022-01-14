package main

import (
	"aws-lambda-standard-replies/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dynaClient, err := handler.DynamodbConfig()
	if err != nil {
		log.Printf("ErrorToLoadDefaultConfig, %s", err)
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDefaultConfig")}), nil
	}
	fmt.Println("req.Resource = ", req.Resource)
	fmt.Println("table = ", os.Getenv("TABLE"))

	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/api/admin/reply/id/{id}":
			return handler.GetStdReplyByID(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/replies":
			return handler.GetAllStdReplies(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/replies/channel/{channel}":
			return handler.GetStdRepliesByChannel(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/replies/name":
			return handler.GetAllRepliesName(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		case "/api/admin/reply":
			return handler.AddStdReply(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/api/admin/reply":
			return handler.UpdateStdReplyByID(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "DELETE":
		switch req.Resource {
		case "/api/admin/reply/id/{id}":
			return handler.DeleteStdReplyByID(req, os.Getenv("TABLE"), dynaClient)

		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	default:
		return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
