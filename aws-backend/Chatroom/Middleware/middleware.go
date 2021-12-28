package Middleware

import (
	"github.com/aws/aws-lambda-go/events"
)

type MiddlwareHandler func(events.APIGatewayProxyRequest) (interface{}, error)

func Middleware(next MiddlwareHandler) MiddlwareHandler {
	return MiddlwareHandler(func(request events.APIGatewayProxyRequest) (interface{}, error) {
		return next(request)
	})
}
