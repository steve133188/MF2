package main

import (
	"aws-lambda-dashboardrequest/handler"
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
	//fmt.Println("table = ", os.Getenv("TABLE"))
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/dashboard":
			return handler.GetDashboard(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
		case "/dashboard/channel":
			return handler.GetDashboardByChannel(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
		case "/dashboard/livechat":
			return handler.GetLivechatDashboard(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
		case "/dashboard/livechat/channel":
			return handler.GetLivechatDashboardByChannel(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
		case "/dashboard/agent":
			return handler.GetAgentDashboard(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
		case "/dashboard/agent/channel":
			return handler.GetAgentDashboardByChannel(req, os.Getenv("DASHBOARDTABLE"), dynaClient)
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
