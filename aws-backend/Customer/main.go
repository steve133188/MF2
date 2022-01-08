package main

import (
	"aws-lambda-customer/handler"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// var dynaClient *dynamodb.Client

func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dynaClient, err := handler.DynamodbConfig()
	if err != nil {
		log.Printf("ErrorToLoadDefaultConfig, %s", err)
		return handler.ApiResponse(http.StatusInternalServerError, handler.ErrMsg{aws.String("ErrorToLoadDefaultConfig")}), nil
	}
	log.Println("req.Resource = ", req.Resource)
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/customers":
			return handler.GetCustomerItems(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customer/{id}":
			return handler.GetCustomerItemByID(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/team/{teamId}":
			return handler.GetCustomersByTeamID(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/tag":
			return handler.GetCustomersByTag(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/group/{group}":
			return handler.GetCustomersByGroup(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/agent":
			return handler.GetCustomersByAgentsID(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/channel":
			return handler.GetCustomersByChannel(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "POST":
		switch req.Resource {
		case "/customer":
			return handler.AddCustomerItem(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "PUT":
		switch req.Resource {
		case "/customer":
			return handler.UpdateCustomerItem(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
			// tag ************************************************************************************************************************
		case "/customer/add-tag":
			return handler.AddTagToCustomer(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customer/del-tag":
			return handler.DeleteCustomerTag(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/add-tag":
			return handler.AddTagToCustomers(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/edit-tags":
			return handler.EditCustomersTag(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/del-tags/{id}":
			return handler.DeleteCustomersTag(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
			//group ************************************************************************************************************************
		case "/customer/group":
			return handler.UpdateGroupToCustomer(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/add-group":
			return handler.AddGroupToCustomers(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/group":
			return handler.UpdateCustomersGroup(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		//case "/customers/del-groups":
		//	return handler.UpdateCustomersGroup(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		//team ************************************************************************************************************************
		case "/customer/team":
			return handler.UpdateCustomerTeam(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/add-teams":
			return handler.AddCustomersTeam(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/teams":
			return handler.EditCustomersTeam(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		//case "/customers/del-teams":
		//	return handler.EditCustomersTeam(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		//agent ************************************************************************************************************************
		case "/customer/add-agent":
			return handler.AddAgentToCustomer(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customer/del-agent":
			return handler.DeleteCustomerAgent(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/add-agent":
			return handler.AddAgentToCustomers(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/edit-agent":
			return handler.EditCustomersAgent(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers/del-agent":
			return handler.DeleteCustomersAgent(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	case "DELETE":
		switch req.Resource {
		case "/customer/{id}":
			return handler.DeleteCustomerItem(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		case "/customers":
			return handler.DeleteCustomers(req, os.Getenv("CUSTOMERTABLE"), dynaClient)
		default:
			log.Println("InvalidRoute")
			return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidRoute")}), nil
		}
	default:
		log.Println("InvalidHTTPMethod")
		return handler.ApiResponse(http.StatusMethodNotAllowed, handler.ErrMsg{aws.String("InvalidHTTPMethod")}), nil
	}

}

func main() {
	lambda.Start(handleRequest)
	// 	middleware.Middleware(
	// 		middleware.JWTHandler(handleRequest),
	// 	),
	// )
}
