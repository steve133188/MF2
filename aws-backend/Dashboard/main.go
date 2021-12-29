package main

import (
	"mf2-aws-dashboard/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() {
	handler.UpdateLivechat()
	handler.UpdateAgent()

}

func main() {
	lambda.Start(handleRequest)
}
