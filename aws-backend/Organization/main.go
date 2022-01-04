package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"aws-lambda-org/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func orgHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		return nil
	})
	if err != nil {
		log.Println("LoadDefaultConfig    ", err)
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDefaultConfig")}), nil
	}

	svc := dynamodb.NewFromConfig(cfg)
	dynaClient := svc
	fmt.Println(req.HTTPMethod)
	fmt.Println(req.Resource)
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/orgs":
			return handler.GetOrgItems(req, os.Getenv("TABLE"), dynaClient)
		case "/org/{id}":
			return handler.GetOrgItemByID(req, os.Getenv("TABLE"), dynaClient)
		case "/org/team":
			return handler.GetTeamName(req, os.Getenv("TABLE"), dynaClient)
		// case "/org/root":
		// 	return handler.GetRootOrg(req, os.Getenv("TABLE"), dynaClient)
		// case "/org/family/{id}":
		// 	return handler.GetFamilyByID(req, os.Getenv("TABLE"), dynaClient)
		default:
			log.Println("InvalidRoute, ", err)
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		case "/org":
			return handler.AddOrgItem(req, os.Getenv("TABLE"), dynaClient)
		default:
			log.Println("InvalidRoute, ", err)
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/org":
			return handler.PutOrgItem(req, os.Getenv("TABLE"), dynaClient)
		default:
			log.Println("InvalidRoute, ", err)
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil

		}
	case "DELETE":
		switch req.Resource {
		case "/org/{id}":
			return handler.DeleteOrgItem(req, os.Getenv("TABLE"), dynaClient)
		default:
			log.Println("InvalidRoute, ", err)
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}

	default:
		log.Println("InvalidHTTPMethod, ", err)
		return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidRoute")}), nil
	}
}

func main() {
	lambda.Start(orgHandler)
	// lambda.Start(
	// 	middleware.Middleware(
	// 		middleware.JWTHandler(orgHandler),
	// 	),
	// )
}
