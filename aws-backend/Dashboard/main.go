package main

import (
	"fmt"
	"mf2-aws-dashboard/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() {
	err := handler.UpdateLivechat()
	if err != nil {
		fmt.Println("ErrorInUpdatingLiveChat, ", err)
	}
	err = handler.UpdateAgent()
	if err != nil {
		fmt.Println("ErrorInUpdatingAgent, ", err)
	}
}

func main() {
	lambda.Start(handleRequest)
}
