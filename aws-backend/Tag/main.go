package main

import (
	"aws-lambda-tag/handler"
	"aws-lambda-tag/middleware"
	"fmt"
	"log"
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
		log.Printf("ErrorToLoadDefaultConfig, %s", err)
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDefaultConfig")}), nil
	}
	fmt.Println("req.Resource = ", req.Resource)
	fmt.Println("table = ", os.Getenv("TABLE"))
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/api/admin/tag/id/{id}":
			return handler.GetTagByID(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/tags":
			return handler.GetAllTags(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		case "/api/admin/tag":
			return handler.AddTag(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/api/admin/tag":
			return handler.UpdateTag(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "DELETE":
		switch req.Resource {
		case "/api/admin/tag/id/{id}":
			return handler.DeleteTagByID(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/tags":
			return handler.DeleteTags(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	default:
		return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}

}

func main() {
	lambda.Start(
		middleware.Middleware(
			middleware.JWTHandler(handleRequest),
		),
	)
}
