package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetCustomerItemByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customerId := req.PathParameters["id"]

	customer := new(model.Customer)

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

	err = attributevalue.UnmarshalMap(out.Item, &customer)
	if err != nil {
		log.Printf("UnmarshalMapError CustomerID = %v, %s", customerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalMapError")}), nil
	}

	users, team, tags, err := FieldHandler(customer.AgentsID, customer.TeamID, customer.TagsID, dynaClient)
	if err != nil {
		fmt.Printf("ErrorFromFieldHandler, %s", err)
		return ApiResponse(http.StatusPartialContent, customer), nil
	}

	fullCustomer := new(model.FullCustomer)

	err = attributevalue.UnmarshalMap(out.Item, fullCustomer)
	if err != nil {
		log.Printf("UnmarshalMapError CustomerID = %v, %s", customerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalMapError")}), nil
	}

	// sout, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
	// 	TableName:        aws.String(os.Getenv("CHATROOM")),
	// 	FilterExpression: aws.String("room_id = :id"),
	// 	ExpressionAttributeValues: map[string]types.AttributeValue{
	// 		":id": &types.AttributeValueMemberN{Value: customerId},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ApiResponse(http.StatusNotFound, ErrMsg{aws.String(err.Error())}), nil
	// }

	// chatroom := make([]model.ChatRoom, 0)

	// err = attributevalue.UnmarshalListOfMaps(sout.Items, &chatroom)
	// if err != nil {
	// 	log.Printf("UnmarshalListOfMaps CustomerID = %v, %s", customerId, err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalListOfMaps")}), nil
	// }

	// for _, v := range chatroom {
	// 	fullCustomer.Channels = append(fullCustomer.Channels, v.Channel)
	// }

	fullCustomer.Agents = users
	fullCustomer.Team = team
	fullCustomer.Tags = tags

	return ApiResponse(http.StatusOK, fullCustomer), nil
}

func GetCustomerItems(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var customers []model.Customer = make([]model.Customer, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName: aws.String(table),
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println("FailedToScanNextPage   ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ScanningError")}), nil
		}

		var pageCustomers []model.Customer = make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pageCustomers)
		if err != nil {
			log.Println("UnmarshalListOfMapsError   ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalListOfMapsError")}), nil
		}

		customers = append(customers, pageCustomers...)

	}
	return ApiResponse(http.StatusOK, customers), nil
}

func GetCustomersByTeamID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	teamId := req.PathParameters["teamId"]

	p := dynamodb.NewQueryPaginator(dynaClient, &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("team_id-index"),
		KeyConditionExpression: aws.String("team_id = :val"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberN{Value: teamId},
		},
	})
	var customers []model.Customer = make([]model.Customer, 0)

	for p.HasMorePages() {
		outs, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("error in query,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in query, " + err.Error())}), nil
		}
		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pCustomers)
		if err != nil {
			fmt.Println("error in unmarshal customer list,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("error in unmarshal customer list, " + err.Error())}), nil
		}
		customers = append(customers, pCustomers...)
	}

	return ApiResponse(http.StatusOK, customers), nil

}

func GetCustomersByGroup(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	group := req.PathParameters["group"]

	var customers []model.Customer = make([]model.Customer, 0)

	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("customer_group = :g"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberS{Value: group},
		},
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemNotExisted: %s", err)
			return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		fmt.Printf("FailedToScan, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &customers)
	if err != nil {
		fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
	}
	return ApiResponse(http.StatusOK, customers), nil

}

func GetCustomersByTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("req.Parameters = ", req.PathParameters)
	fmt.Println("taglist = ", req.MultiValueQueryStringParameters["tag"])

	tagIdList := req.MultiValueQueryStringParameters["tag"]
	// "contains(tags_id, :t1) AND contains(tags_id, :t2)"
	dataVal := make(map[string]types.AttributeValue)
	var filterStr string

	for k, v := range tagIdList {
		if k != 0 {
			filterStr += " AND contains(tags_id, :t" + strconv.Itoa(k) + ")"
		} else {
			filterStr += "contains(tags_id, :t" + strconv.Itoa(k) + ")"
		}

		key := ":t" + strconv.Itoa(k)
		dataVal[key] = &types.AttributeValueMemberN{Value: v}

	}

	res, _ := json.Marshal(dataVal)

	fmt.Println("dataVal = ", string(res))
	fmt.Println("filterStr = ", filterStr)

	var customers []model.Customer = make([]model.Customer, 0)

	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(table),
		FilterExpression:          aws.String(filterStr),
		ExpressionAttributeValues: dataVal,
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemNotExisted: %s", err)
			return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		fmt.Printf("FailedToScan, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
	}

	fmt.Println("count = ", out.Count)

	err = attributevalue.UnmarshalListOfMaps(out.Items, &customers)
	if err != nil {
		fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
	}

	//fullCustomers := make([]model.FullCustomer, 0)
	//
	//for k, v := range customers {
	//	fullCustomer := new(model.FullCustomer)
	//	users, team, tags, err := FieldHandler(v.AgentsID, v.TeamID, v.TagsID, dynaClient)
	//	if err != nil {
	//		fmt.Printf("ErrorFromFieldHandler, %s", err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorFromFieldHandler")}), nil
	//	}
	//
	//	err = attributevalue.UnmarshalMap(out.Items[k], &fullCustomer)
	//	if err != nil {
	//		fmt.Printf("FailedToUnmarshalMap, %s", err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	//	}
	//	fmt.Printf("fullcustomer = %v", fullCustomer)
	//	fullCustomer.Agents = users
	//	fullCustomer.Team = team
	//	fullCustomer.Tags = tags
	//
	//	sout, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
	//		TableName:        aws.String(os.Getenv("CHATROOM")),
	//		FilterExpression: aws.String("room_id = :id"),
	//		ExpressionAttributeValues: map[string]types.AttributeValue{
	//			":id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
	//		},
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	//	}
	//
	//	chatroom := make([]model.ChatRoom, 0)
	//
	//	err = attributevalue.UnmarshalListOfMaps(sout.Items, &chatroom)
	//	if err != nil {
	//		log.Printf("UnmarshalListOfMaps CustomerID = %v, %s", strconv.Itoa(v.CustomerID), err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalListOfMaps")}), nil
	//	}
	//
	//	for _, v := range chatroom {
	//		fullCustomer.Channels = append(fullCustomer.Channels, v.Channel)
	//	}
	//	fullCustomers = append(fullCustomers, *fullCustomer)
	//}

	return ApiResponse(http.StatusOK, customers), nil

}

func GetCustomersByAgentID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	agentId := req.PathParameters["agent"]

	p := dynamodb.NewQueryPaginator(dynaClient, &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("handler_id-index"),
		KeyConditionExpression: aws.String("handler_id = :hid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hid": &types.AttributeValueMemberN{Value: agentId},
		},
	})

	customers := make([]model.Customer, 0)
	for p.HasMorePages() {
		outs, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("Error in query,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("Error in query, " + err.Error())}), nil
		}

		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pCustomers)
		if err != nil {
			fmt.Println("Error in unmarshal list of data,", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("Error in unmarshal list of data, " + err.Error())}), nil
		}

		customers = append(customers, pCustomers...)
	}

	return ApiResponse(http.StatusOK, customers), nil
}

func GetCustomersByChannel(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("req.Parameters = ", req.PathParameters)

	channelsList := req.MultiValueQueryStringParameters["channel"]
	// "contains(tags_id, :t1) AND contains(tags_id, :t2)"
	dataVal := make(map[string]types.AttributeValue)
	var filterStr string

	for k, v := range channelsList {
		if k != 0 {
			filterStr += " AND contains(channels, :t" + strconv.Itoa(k) + ")"
		} else {
			filterStr += "contains(channels, :t" + strconv.Itoa(k) + ")"
		}

		key := ":t" + strconv.Itoa(k)
		dataVal[key] = &types.AttributeValueMemberS{Value: v}

	}

	res, _ := json.Marshal(dataVal)

	fmt.Println("dataVal = ", string(res))
	fmt.Println("filterStr = ", filterStr)

	var customers []model.Customer = make([]model.Customer, 0)

	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(table),
		FilterExpression:          aws.String(filterStr),
		ExpressionAttributeValues: dataVal,
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemNotExisted: %s", err)
			return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		fmt.Printf("FailedToScan, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan")}), nil
	}

	fmt.Println("count = ", out.Count)

	err = attributevalue.UnmarshalListOfMaps(out.Items, &customers)
	if err != nil {
		fmt.Printf("FailedToUnmarshalListOfMap, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMap")}), nil
	}

	//fullCustomers := make([]model.FullCustomer, 0)
	//
	//for k, v := range customers {
	//	fullCustomer := new(model.FullCustomer)
	//	users, team, tags, err := FieldHandler(v.AgentsID, v.TeamID, v.TagsID, dynaClient)
	//	if err != nil {
	//		fmt.Printf("ErrorFromFieldHandler, %s", err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorFromFieldHandler")}), nil
	//	}
	//
	//	err = attributevalue.UnmarshalMap(out.Items[k], &fullCustomer)
	//	if err != nil {
	//		fmt.Printf("FailedToUnmarshalMap, %s", err)
	//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	//	}
	//	fmt.Printf("fullcustomer = %v", fullCustomer)
	//	fullCustomer.Agents = users
	//	fullCustomer.Team = team
	//	fullCustomer.Tags = tags
	//
	//	fullCustomers = append(fullCustomers, *fullCustomer)
	//}

	return ApiResponse(http.StatusOK, customers), nil

}
