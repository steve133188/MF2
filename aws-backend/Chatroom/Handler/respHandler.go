package Handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ApiResponse(status int, body interface{}) *events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: status,
	}

	byteBody, _ := json.Marshal(body)

	resp.Body = string(byteBody)
	return &resp
}
