package main

import (
	"aws-lambda-user/handler"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dynaClient, err := handler.DynamodbConfig()
	if err != nil {
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDynamodbConfig")}), nil
	}

	table := os.Getenv("TABLE")
	fmt.Println("req.HTTPMethod, ", req.HTTPMethod)
	fmt.Println("req.Resource, ", req.Resource)

	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/api/users/{id}":
			return handler.GetUserByID(req, table, dynaClient)
		case "/api/users/team/{teamId}":
			return handler.GetUsersByTeamID(req, table, dynaClient)
		case "/api/users/role/{roleId}":
			return handler.GetUsersByRoleID(req, table, dynaClient)
		case "/api/users/all":
			return handler.GetUsers(req, table, dynaClient)
		case "/api/users/no-team":
			return handler.GetUsersWithoutTeam(req, table, dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidGetMethod")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/api/users":
			return handler.UpdateUser(req, table, dynaClient)
		case "/api/users/status":
			return handler.UpdateUserStatus(req, table, dynaClient)
		case "/api/users/change-password":
			return handler.UpdateUserPassword(req, table, dynaClient)
		case "/api/user/team":
			return handler.UpdateUserTeam(req, table, dynaClient)
		case "/api/users/team":
			return handler.UpdateUsersTeam(req, table, dynaClient)
		case "/api/user/role":
			return handler.UpdateUserRole(req, table, dynaClient)
		case "/api/users/role":
			return handler.UpdateUsersRole(req, table, dynaClient)
		case "/api/user/add-channels":
			return handler.AddUserChannels(req, table, dynaClient)
		case "/api/user/edit-channels":
			return handler.EditUserChannels(req, table, dynaClient)
		case "/api/user/del-channels":
			return handler.DeleteUserChannels(req, table, dynaClient)
		case "/api/user/access-right":
			return handler.UpdateUserChatAccess(req, table, dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidPutMethod")}), nil

		}
	case "DELETE":
		switch req.Resource {
		case "/api/user/{id}":
			return handler.DeleteUserByID(req, table, dynaClient)
		case "/api/users":
			return handler.DeleteUsers(req, table, dynaClient)
		default:
			return handler.ApiResponse(http.StatusBadRequest, handler.ErrMsg{aws.String("InvalidDeleteMethod")}), nil

		}

	default:
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}

}

func main() {
	// lambda.Start(
	// 	middleware.Middleware(
	// 		middleware.JWTHandler(handleRequest),
	// 	),
	// )
	lambda.Start(handleRequest)
}
