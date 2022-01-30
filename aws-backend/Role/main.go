package main

import (
	"aws-lambda-role/handler"
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
		case "/api/admin/role/id/{id}":
			return handler.GetRoleByID(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/roles":
			return handler.GetAllRoles(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		case "/api/admin/role":
			return handler.AddRole(req, os.Getenv("TABLE"), dynaClient)
		case "/api/admin/roles/delete":
			return handler.DeleteRoles(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/api/admin/role":
			return handler.UpdateRoleWithID(req, os.Getenv("TABLE"), dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "DELETE":
		switch req.Resource {
		case "/api/admin/role/id/{id}":
			return handler.DeleteRoleByID(req, os.Getenv("TABLE"), dynaClient)

		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	default:
		return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}

}

func main() {
	lambda.Start(handleRequest)
	// lambda.Start(
	// 	middleware.Middleware(
	// 		middleware.JWTHandler(handleRequest),
	// 	),
	// )
}
