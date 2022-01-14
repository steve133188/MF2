package handler

import (
	"aws-lambda-customer/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateCustomerItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customer := new(model.Customer)

	err := json.Unmarshal([]byte(req.Body), &customer)
	if err != nil {
		log.Printf("FailedToUnmarshalInputData: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalInputData")}), nil
	}

	customer.UpdateAt = time.Now().Unix()

	av, err := attributevalue.MarshalMap(&customer)
	if err != nil {
		log.Println("EditItem MarshalMap    ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMapData")}), nil

	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Printf("ItemNotExisted: %s", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		log.Printf("ErrorToUpdateItem: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorToUpdateItem")}), nil

	}

	return ApiResponse(http.StatusOK, nil), nil
}

//customer team ************************************************************************************************************************

func UpdateCustomerTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customer := new(model.Customer)

	err := json.Unmarshal([]byte(req.Body), &customer)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{(aws.String("FailedToUnmarshalReqBody"))}), nil
	}

	customer.UpdateAt = time.Now().Unix()

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(customer.CustomerID)},
		},
		UpdateExpression: aws.String("set team_id = :te, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":te": &types.AttributeValueMemberS{Value: strconv.Itoa(customer.TeamID)},
			":t":  &types.AttributeValueMemberS{Value: strconv.FormatInt(customer.UpdateAt, 10)},
		},
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
	})
	if err != nil {
		fmt.Println("FailedToUpdate, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdate")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func AddCustomersTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customerIDs := req.MultiValueQueryStringParameters["id"]
	var data struct {
		Team string `json:"team"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	time := time.Now().Unix()

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET customer_group = :g, update_at = :t"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":g": &types.AttributeValueMemberS{Value: data.Team},
				":t": &types.AttributeValueMemberN{Value: strconv.FormatInt(time, 10)},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToAddTeamToMany, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToAddTeamToMany, CustomerID = " + v + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

func EditCustomersTeam(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		OldTeam string `json:"old_team"`
		NewTeam string `json:"new_team"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("team_id = :g"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberN{Value: data.OldTeam},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	time := time.Now().Unix()

	for _, v := range customers {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("Set #cg = :ng, update_at = :t"),
			ExpressionAttributeNames: map[string]string{
				"#cg": "team_id",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: data.NewTeam},
				":t":  &types.AttributeValueMemberN{Value: strconv.FormatInt(time, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateTeam, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateTeam, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

//func DeleteCustomerTeams(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
//	var data struct {
//		DeleteTeam string `json:"delete_team"`
//	}
//
//	err := json.Unmarshal([]byte(req.Body), &data)
//	if err != nil {
//		fmt.Println("FailedToUnmarshalReqData, ", err)
//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
//	}
//
//	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
//		TableName:        &table,
//		FilterExpression: aws.String("team_id = :g"),
//		ExpressionAttributeValues: map[string]types.AttributeValue{
//			":g": &types.AttributeValueMemberN{Value: data.DeleteTeam},
//		},
//		Limit: aws.Int32(50),
//	})
//
//	customers := make([]model.Customer, 0)
//
//	for p.HasMorePages() {
//		out, err := p.NextPage(context.TODO())
//		if err != nil {
//			fmt.Println("ErrorInNextPage, ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
//		}
//
//		pItems := make([]model.Customer, 0)
//		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
//		if err != nil {
//			fmt.Println("UnmarshalListOfMaps, ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
//		}
//
//		customers = append(customers, pItems...)
//	}
//
//	for _, v := range customers {
//		removeStr := "remove team_id"
//		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
//			TableName: aws.String(table),
//			Key: map[string]types.AttributeValue{
//				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
//			},
//			UpdateExpression: aws.String(removeStr),
//		})
//		if err != nil {
//			fmt.Println("FailedToDeleteTeam, CustomerID = ", v.CustomerID, ", ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteTeam, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
//		}
//	}
//
//	return ApiResponse(http.StatusOK, nil), nil
//
//}

//customer group ************************************************************************************************************************

func UpdateGroupToCustomer(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customer := new(model.Customer)

	err := json.Unmarshal([]byte(req.Body), &customer)
	if err != nil {
		fmt.Printf("FailedToUnmarshalReqBody, %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{(aws.String("FailedToUnmarshalReqBody"))}), nil
	}

	customer.UpdateAt = time.Now().Unix()

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(customer.CustomerID)},
		},
		UpdateExpression: aws.String("set customer_group = :g, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberS{Value: customer.Group},
			":t": &types.AttributeValueMemberN{Value: strconv.FormatInt(customer.UpdateAt, 10)},
		},
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
	})
	if err != nil {
		fmt.Println("FailedToUpdate, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdate")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func UpdateCustomersGroup(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		OldGroup string `json:"old_group"`
		NewGroup string `json:"new_group"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("customer_group = :g"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberS{Value: data.OldGroup},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	time := time.Now().Unix()

	for _, v := range customers {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("Set #cg = :ng, update_at = :t"),
			ExpressionAttributeNames: map[string]string{
				"#cg": "customer_group",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberS{Value: data.NewGroup},
				":t":  &types.AttributeValueMemberN{Value: strconv.FormatInt(time, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateItem, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

func AddGroupToCustomers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customerIDs := req.MultiValueQueryStringParameters["id"]
	var data struct {
		Group string `json:"group"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	time := time.Now().Unix()

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET customer_group = :g, update_at = :t"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":g": &types.AttributeValueMemberS{Value: data.Group},
				":t": &types.AttributeValueMemberN{Value: strconv.FormatInt(time, 10)},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToAddTagToMany, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToAddTagToMany, CustomerID = " + v + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

//func DeleteCustomersGroup(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
//	var data struct {
//		DeleteGroup string `json:"delete_group"`
//	}
//
//	err := json.Unmarshal([]byte(req.Body), &data)
//	if err != nil {
//		fmt.Println("FailedToUnmarshalReqData, ", err)
//		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
//	}
//
//	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
//		TableName:        &table,
//		FilterExpression: aws.String("customer_group = :g"),
//		ExpressionAttributeValues: map[string]types.AttributeValue{
//			":g": &types.AttributeValueMemberS{Value: data.DeleteGroup},
//		},
//		Limit: aws.Int32(50),
//	})
//
//	customers := make([]model.Customer, 0)
//
//	for p.HasMorePages() {
//		out, err := p.NextPage(context.TODO())
//		if err != nil {
//			fmt.Println("ErrorInNextPage, ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
//		}
//
//		pItems := make([]model.Customer, 0)
//		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
//		if err != nil {
//			fmt.Println("UnmarshalListOfMaps, ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
//		}
//
//		customers = append(customers, pItems...)
//	}
//
//	for _, v := range customers {
//		removeStr := "remove customer_group"
//		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
//			TableName: aws.String(table),
//			Key: map[string]types.AttributeValue{
//				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
//			},
//			UpdateExpression: aws.String(removeStr),
//		})
//		if err != nil {
//			fmt.Println("FailedToDeleteGroup, CustomerID = ", v.CustomerID, ", ", err)
//			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteGroup, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
//		}
//	}
//
//	return ApiResponse(http.StatusOK, nil), nil
//
//}

//customer tag ************************************************************************************************************************

func AddTagToCustomer(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		CustomerId int   `json:"customer_id"`
		Tags       []int `json:"tag_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	check, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	if check.Item == nil {
		fmt.Println("CustomerIDNotExisted, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("CustomerIDNotExisted")}), nil
	}

	customer := new(model.Customer)
	err = attributevalue.UnmarshalMap(check.Item, &customer)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	for _, v := range data.Tags {
		tagCheck, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("TAGTABLE")),
			Key: map[string]types.AttributeValue{
				"tag_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetItem, Tag_id = ", v, err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetTagItem")}), nil
		}
		if tagCheck.Item == nil {
			fmt.Println("TagNotExist, tagID = ", v)
			return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("TagNotExisted, tagID = " + strconv.Itoa(v))}), nil
		}

		for _, j := range customer.TagsID {
			if v == j {
				fmt.Println("TagNotExist, tagID = ", v, err)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("TagExisted, tagID = " + strconv.Itoa(v))}), nil
			}
		}
	}

	res, err := attributevalue.MarshalList(data.Tags)
	if err != nil {
		fmt.Println("FailedToMarshalList, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalList")}), nil
	}

	updateTime := time.Now().Unix()

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression: aws.String("SET tags_id = list_append(tags_id, :t), update_at = :u"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberL{Value: res},
			":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	fmt.Println("UpdateResult, ", out.Attributes)

	return ApiResponse(http.StatusOK, nil), nil
}

func AddTagToCustomers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	//customerIDs := req.MultiValueQueryStringParameters["id"]
	var data struct {
		CustomerID []int `json:"customer_id"`
		Tags       []int `json:"tags_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	res, err := attributevalue.MarshalList(data.Tags)
	if err != nil {
		fmt.Println("FailedToMarshalList, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalList")}), nil
	}

	updateTime := time.Now().Unix()

	for _, v := range data.CustomerID {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
			UpdateExpression: aws.String("SET tags_id = list_append(tags_id, :t), update_at = :u "),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":t": &types.AttributeValueMemberL{Value: res},
				":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToUpdateTag, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateTag, CustomerID = " + strconv.Itoa(v) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func EditCustomersTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		OldTag string `json:"old_tag"`
		NewTag string `json:"new_tag"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("contains(tags_id, :g)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberN{Value: data.OldTag},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	// change old tag to int
	oldTagInt, _ := strconv.Atoi(data.OldTag)

	for _, v := range customers {

		//set index for update tag
		var tagIndex string
		for i, old := range v.TagsID {
			if old == oldTagInt {
				tagIndex = strconv.Itoa(i)
				break
			}
		}

		updateTime := time.Now().Unix()

		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("SET tags_id[" + tagIndex + "].field = :ng, update_at = :u"),
			ReturnValues:     types.ReturnValueUpdatedNew,
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: data.NewTag},
				":u":  &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateTag, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateTag, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

func DeleteCustomerTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		CustomerId int   `json:"customer_id"`
		Tags       []int `json:"tag_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	cout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetCustomer, CustomerID = ", data.CustomerId)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetCustomer, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}
	if cout.Item == nil {
		fmt.Println("CustomerNotExisted, CustomerID = ", data.CustomerId)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("CustomerNotExisted, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}

	customer := new(model.Customer)
	err = attributevalue.UnmarshalMap(cout.Item, &customer)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, CustomerID = ", data.CustomerId)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}

	// var tagStr string
	fmt.Println("data.tags = ", data.Tags)
	fmt.Println("customer.TagsID = ", customer.TagsID)
	for k, v := range customer.TagsID {
		for _, j := range data.Tags {
			fmt.Println("customer.TagsID = ", k)
			fmt.Println("data.tags = ", j)

			if v == j {
				gout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
					TableName: aws.String(os.Getenv("TAGTABLE")),
					Key: map[string]types.AttributeValue{
						"tag_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
					},
				})
				if err != nil {
					fmt.Println("FailedToGetTag, Tag_id = ", v, err)
					return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetTag, tagID = " + strconv.Itoa(v))}), nil
				}
				if gout.Item == nil {
					fmt.Println("TagNotExist, tag_id = ", v)
					return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("TagNotExist, tagID = " + strconv.Itoa(v))}), nil
				}

				customer.TagsID = remove(customer.TagsID, k)
			}
		}

	}

	tagList, err := attributevalue.MarshalList(customer.TagsID)
	if err != nil {
		fmt.Println("FailedToMarshaList CustomerID = ", data.CustomerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshaList CustomerID = " + strconv.Itoa(data.CustomerId) + err.Error())}), nil
	}

	// fmt.Println("tagStr = ", tagStr)
	updateTime := time.Now().Unix()

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression:    aws.String("Set update_at = :u, tags_id = :tid"),
		ReturnValues:        types.ReturnValueUpdatedNew,
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":u":   &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			":tid": &types.AttributeValueMemberL{Value: tagList},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func DeleteCustomersTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	// var data struct {
	// 	DeleteTag string `json:"delete_tag"`
	// }

	id := req.PathParameters["id"]

	// err := json.Unmarshal([]byte(req.Body), &data)
	// if err != nil {
	// 	fmt.Println("FailedToUnmarshalReqData, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	// }

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("contains(tags_id, :g)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberN{Value: id},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	// change old tag to int
	deleteTagInt, _ := strconv.Atoi(id)

	for _, v := range customers {

		//set update index
		var tagIndex string
		for i, old := range v.TagsID {
			if old == deleteTagInt {
				tagIndex = strconv.Itoa(i)
				break
			}
		}

		updateTime := time.Now().Unix()

		removeStr := "remove tags_id[" + tagIndex + "]"
		_, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String(removeStr + " Set update_at = :u"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToRemoveTag, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToRemoveTag, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}

//customer agent ************************************************************************************************************************

func AddAgentToCustomer(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		CustomerId int   `json:"customer_id"`
		Agents     []int `json:"agents_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	check, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem")}), nil
	}

	if check.Item == nil {
		fmt.Println("CustomerIDNotExisted, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("CustomerIDNotExisted, customerID = " + strconv.Itoa(data.CustomerId))}), nil
	}

	customer := new(model.Customer)
	err = attributevalue.UnmarshalMap(check.Item, &customer)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap")}), nil
	}

	for _, v := range data.Agents {
		userCheck, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("USERTABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetItem, user_id = ", v, err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetTagItem, user_id = " + strconv.Itoa(v))}), nil
		}
		if userCheck.Item == nil {
			fmt.Println("UserNotExist, user_id = ", v)
			return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("UserNotExist, user_id = " + strconv.Itoa(v))}), nil
		}

		for _, j := range customer.AgentsID {
			if v == j {
				fmt.Println("UserIDExisted, user_id = ", v)
				return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UserIDExisted, user_id = " + strconv.Itoa(v))}), nil
			}
		}
	}

	res, err := attributevalue.MarshalList(data.Agents)
	if err != nil {
		fmt.Println("FailedToMarshalList, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalList")}), nil
	}

	updateTime := time.Now().Unix()

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression: aws.String("SET agents_id = list_append(agents_id, :t), update_at = :u"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberL{Value: res},
			":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	fmt.Println("UpdateResult, ", out.Attributes)
	// err = ChangeAgentLeads('+', 1, data.Agents, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, nil), nil
}

func AddAgentToCustomers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		CustomerID []int `json:"customer_id"`
		AgentID    []int `json:"agents_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	res, err := attributevalue.MarshalList(data.AgentID)
	if err != nil {
		fmt.Println("FailedToMarshalList, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalList")}), nil
	}

	updateTime := time.Now().Unix()

	for _, v := range data.CustomerID {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
			UpdateExpression: aws.String("SET agents_id = list_append(agents_id, :t), update_at = :u"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":t": &types.AttributeValueMemberL{Value: res},
				":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToAddAgentToMany, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToAddAgentToMany, CustomerID = " + strconv.Itoa(v) + ", " + err.Error())}), nil
		}
	}

	// err = ChangeAgentLeads('+', len(data.CustomerID), data.AgentID, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, nil), nil
}

func EditCustomersAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		OldAgent int `json:"old_agent"`
		NewAgent int `json:"new_agent"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("contains(agents_id, :g)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberN{Value: strconv.Itoa(data.OldAgent)},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	// change old tag to int

	for _, v := range customers {

		//set index for update tag
		var agentIndex string
		for i, old := range v.AgentsID {
			if old == data.OldAgent {
				agentIndex = strconv.Itoa(i)
				break
			}
		}
		updateTime := time.Now().Unix()
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("set agents_id[" + agentIndex + "] = :ng update_at = :u"),
			ReturnValues:     types.ReturnValueUpdatedNew,
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: strconv.Itoa(data.NewAgent)},
				":u":  &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateAgentToMany, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateAgentToMany, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	// tempList := make([]int, 0)
	// tempList = append(tempList, data.OldAgent)
	// err = ChangeAgentLeads('-', len(customers), tempList, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }
	// tempList[0] = data.NewAgent
	// err = ChangeAgentLeads('+', len(customers), tempList, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, nil), nil

}

func DeleteCustomerAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		CustomerId int   `json:"customer_id"`
		Agents     []int `json:"agents_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	customer := new(model.Customer)

	cout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetCustomer, CustomerID = ", data.CustomerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetCustomer, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}
	if cout.Item == nil {
		fmt.Println("CustomerNotExist, CustomerID = ", data.CustomerId)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("CustoemrNotExist, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}

	err = attributevalue.UnmarshalMap(cout.Item, &customer)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, CustomerID = ", data.CustomerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, CustomerID = " + strconv.Itoa(data.CustomerId))}), nil
	}

	// var tagStr string

	for _, v := range data.Agents {
		gout, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(os.Getenv("USERTABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetAgent, userID = ", v, err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetAgent, userID = " + strconv.Itoa(v))}), nil
		}
		if gout.Item == nil {
			fmt.Println("AgentNotExist, userID = ", v)
			return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("AgentNotExist, userID = " + strconv.Itoa(v))}), nil
		}

		for i, j := range customer.AgentsID {
			if j == v {
				customer.AgentsID = remove(customer.AgentsID, i)
			}
		}

	}

	a_list, err := attributevalue.MarshalList(customer.AgentsID)
	if err != nil {
		fmt.Println("FailedToMarshalList, CustomerId = ", data.CustomerId, err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String(err.Error())}), nil
	}

	updateTime := time.Now().Unix()

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression:    aws.String("SET update_at = :u, agents_id = :a"),
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
		ReturnValues:        types.ReturnValueUpdatedNew,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			":a": &types.AttributeValueMemberL{Value: a_list},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	// err = ChangeAgentLeads('-', 1, data.Agents, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, nil), nil
}

func DeleteCustomersAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		DeleteAgent int `json:"delete_agent"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqData, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqData")}), nil
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		FilterExpression: aws.String("contains(agents_id, :g)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberN{Value: strconv.Itoa(data.DeleteAgent)},
		},
		Limit: aws.Int32(50),
	})

	customers := make([]model.Customer, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}

		pItems := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		customers = append(customers, pItems...)
	}

	// change old tag to int
	updateTime := time.Now().Unix()
	for _, v := range customers {

		//set update index
		var agentIndex string
		for i, old := range v.AgentsID {
			if old == data.DeleteAgent {
				agentIndex = strconv.Itoa(i)
				break
			}
		}
		fmt.Println("Agent ID found in Customer:", v.CustomerID, " Index = ", agentIndex)

		removeStr := "REMOVE agents_id[" + agentIndex + "]"
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String(removeStr + " SET update_at = :u"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":u": &types.AttributeValueMemberN{Value: strconv.FormatInt(updateTime, 10)},
			},
		})
		if err != nil {
			fmt.Println("FailedToRemoveAgent, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToRemoveAgent, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	// tempList := make([]int, 0)
	// tempList = append(tempList, data.DeleteAgent)
	// err = ChangeAgentLeads('-', len(customers), tempList, dynaClient)
	// if err != nil {
	// 	fmt.Println("FailedToChangeLeads, ", err)
	// 	return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToChangeLeads")}), nil

	// }

	return ApiResponse(http.StatusOK, nil), nil

}

// func ChangeAgentLeads(operator byte, changeValue int, agentID []int, dynaClient *dynamodb.Client) error {

// 	dataVal := make(map[string]types.AttributeValue)
// 	var filterStr string

// 	for k, v := range agentID {
// 		if k != 0 {
// 			filterStr += " OR user_id = :t" + strconv.Itoa(k)
// 		} else {
// 			filterStr += "user_id = :t" + strconv.Itoa(k)
// 		}

// 		key := ":t" + strconv.Itoa(k)
// 		dataVal[key] = &types.AttributeValueMemberN{Value: strconv.Itoa(v)}

// 	}

// 	agents := make([]model.User, 0)
// 	out, err := dynaClient.Scan(context.TODO(), &dynamodb.ScanInput{
// 		TableName:                 aws.String(os.Getenv("USERTABLE")),
// 		FilterExpression:          aws.String(filterStr),
// 		ExpressionAttributeValues: dataVal,
// 	})
// 	if err != nil {
// 		fmt.Printf("FailedToScan User Leads, %s", err)
// 		return err
// 	}

// 	fmt.Println("count = ", out.Count)

// 	err = attributevalue.UnmarshalListOfMaps(out.Items, &agents)
// 	if err != nil {
// 		fmt.Printf("FailedToUnmarshalListOfMap User Leads, %s", err)
// 		return err
// 	}

// 	for _, v := range agents {

// 		//set update index
// 		switch operator {
// 		case '+':
// 			v.Leads += changeValue
// 		case '-':
// 			v.Leads -= changeValue
// 		default:
// 			break
// 		}

// 		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
// 			TableName: aws.String(os.Getenv("USERTABLE")),
// 			Key: map[string]types.AttributeValue{
// 				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.UserID)},
// 			},
// 			UpdateExpression: aws.String("SET leads = :l"),
// 			ExpressionAttributeValues: map[string]types.AttributeValue{
// 				":l": &types.AttributeValueMemberN{Value: strconv.Itoa(v.Leads)},
// 			},
// 			ReturnValues: types.ReturnValueAllNew,
// 		})
// 		if err != nil {
// 			fmt.Println("FailedToChangeLeads, CustomerID = ", v, ", ", err)
// 			return err
// 		}
// 	}

// 	return nil
// }

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
