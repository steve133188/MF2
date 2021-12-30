package main

import (
	"fmt"
	"mf2-aws-dashboard/handler"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("req.Path = ", req.Path)
	fmt.Println("req.resource = ", req.Resource)

	switch req.HTTPMethod {
	case "GET":
		err := handler.UpdateLivechat()
		if err != nil {
			fmt.Println("ErrorInUpdatingLiveChat, ", err)
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "ErrorInUpdatingLiveChat",
			}, err
		}
		err = handler.UpdateAgent()
		if err != nil {
			fmt.Println("ErrorInUpdatingAgent, ", err)
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "ErrorInUpdatingAgent",
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
