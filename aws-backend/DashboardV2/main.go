package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mf2-lambda-dashboardv2/handler"
	"net/http"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("req.Path = ", req.Path)
	fmt.Println("req.resource = ", req.Resource)

	switch req.HTTPMethod {
	case "GET":
		err := handler.UpdateDashBoard()
		if err != nil {
			fmt.Println("ErrorInUpdatingLiveChat, ", err)
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "ErrorInUpdatingLiveChat",
			}, err
		}
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "success",
		}, nil
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "InvalidHTTPRequuest",
		}, nil
	}

}

func main() {
	lambda.Start(handleRequest)
}
