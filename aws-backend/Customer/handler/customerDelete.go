package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteCustomerItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customerId := req.PathParameters["id"]

	if len(customerId) == 0 {
		log.Println("MissingCustomerID")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingCustomerID")}), nil
	}

	//find agent id
	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: customerId},
		},
	})
	if err != nil {
		log.Printf("FailedToGetItem CustomerID = %v, %s", customerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	if out.Item == nil {
		log.Printf("ItemNotExist CustomerID = %v, %s", customerId, err)
		return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotExist")}), nil
	}
	customer := new(model.Customer)
	err = attributevalue.UnmarshalMap(out.Item, &customer)
	if err != nil {
		log.Printf("UnmarshalMapError CustomerID = %v, %s", customerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalMapError")}), nil
	}

	//delete
	out2, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: customerId},
		},
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
		ReturnValues:        types.ReturnValueAllOld,
	})
	if err != nil {

		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemNotExisted: %s", err)
			return ApiResponse(http.StatusGone, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		log.Printf("FailedToDeleteItem: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem")}), nil
	}

	log.Println("out.Attributes = ", out2.Attributes)

	// err = ChangeAgentLeads('-', 1, customer.AgentsID, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, map[string]string{"message": "success"}), nil
}

func DeleteCustomers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {

	var data struct {
		CustomerID []string `json:"customer_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	for _, v := range data.CustomerID {
		//find agent id
		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
		})
		if err != nil {
			log.Printf("FailedToGetItem CustomerID = %v, %s", v, err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
		}

		if out.Item == nil {
			log.Printf("ItemNotExist CustomerID = %v, %s", v, err)
			return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotExist")}), nil
		}
		customer := new(model.Customer)
		err = attributevalue.UnmarshalMap(out.Item, &customer)
		if err != nil {
			log.Printf("UnmarshalMapError CustomerID = %v, %s", v, err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalMapError")}), nil
		}
		// err = ChangeAgentLeads('-', 1, customer.AgentsID, dynaClient)
		// if err != nil {
		// 	fmt.Println("FailedToChangeLeads, ", err)
		// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

		// }

		//delete
		_, err = dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
		})
		if err != nil {
			// fmt.Println("FailedToDeleteItem, UserID = ", v, ", ", err)
			fmt.Println(err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem, CustomerID = " + v + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil
}
