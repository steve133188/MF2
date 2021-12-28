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

	customer.UpdateAt = time.Now().Format(os.Getenv("TIMEFORMAT"))

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

	customer.UpdateAt = time.Now().Format(os.Getenv("TIMEFORMAT"))

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(customer.CustomerID)},
		},
		UpdateExpression: aws.String("set team_id = :te, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":te": &types.AttributeValueMemberS{Value: strconv.Itoa(customer.TeamID)},
			":t":  &types.AttributeValueMemberS{Value: customer.UpdateAt},
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

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET customer_group = :g"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":g": &types.AttributeValueMemberS{Value: data.Team},
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

	for _, v := range customers {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("Set #cg = :ng"),
			ExpressionAttributeNames: map[string]string{
				"#cg": "team_id",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: data.NewTeam},
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

	customer.UpdateAt = time.Now().Format(os.Getenv("TIMEFORMAT"))

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(customer.CustomerID)},
		},
		UpdateExpression: aws.String("set customer_group = :g, update_at = :t"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":g": &types.AttributeValueMemberS{Value: customer.Group},
			":t": &types.AttributeValueMemberS{Value: customer.UpdateAt},
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

	for _, v := range customers {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("Set #cg = :ng"),
			ExpressionAttributeNames: map[string]string{
				"#cg": "customer_group",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberS{Value: data.NewGroup},
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

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET customer_group = :g"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":g": &types.AttributeValueMemberS{Value: data.Group},
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
		Tags       []int `json:"tags_id"`
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

	updateTime := time.Now().Format(os.Getenv("TIMEFORMAT"))

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression: aws.String("SET tags_id = list_append(tags_id, :t), update_at = :u"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberL{Value: res},
			":u": &types.AttributeValueMemberS{Value: updateTime},
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
	customerIDs := req.MultiValueQueryStringParameters["id"]
	var data struct {
		Tags []int `json:"tags_id"`
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

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET tags_id = list_append(tags_id, :t)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":t": &types.AttributeValueMemberL{Value: res},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToUpdateTag, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateTag, CustomerID = " + v + ", " + err.Error())}), nil
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

		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("set tags_id[" + tagIndex + "].field = :ng"),
			ReturnValues:     types.ReturnValueUpdatedNew,
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: data.NewTag},
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
		Tags       []int `json:"tags_id"`
	}

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("FailedToUnmarshalReqBody, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalReqBody")}), nil
	}

	var tagStr string

	for k, v := range data.Tags {
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

		if k == 0 {
			tagStr = "REMOVE tags_id[" + strconv.Itoa(k) + "]"
		} else {
			tagStr += ", tags_id[" + strconv.Itoa(k) + "]"
		}

	}

	updateTime := time.Now().Format(os.Getenv("TIMEFORMAT"))

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression:    aws.String(tagStr + ", Set update_at = :u"),
		ReturnValues:        types.ReturnValueUpdatedNew,
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":u": &types.AttributeValueMemberS{Value: updateTime},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func DeleteCustomersTag(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		DeleteTag string `json:"delete_tag"`
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
			":g": &types.AttributeValueMemberN{Value: data.DeleteTag},
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
	deleteTagInt, _ := strconv.Atoi(data.DeleteTag)

	for _, v := range customers {

		//set update index
		var tagIndex string
		for i, old := range v.TagsID {
			if old == deleteTagInt {
				tagIndex = strconv.Itoa(i)
				break
			}
		}

		removeStr := "remove tags_id[" + tagIndex + "]"
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String(removeStr),
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

	updateTime := time.Now().Format(os.Getenv("TIMEFORMAT"))

	out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression: aws.String("SET agents_id = list_append(agents_id, :t), update_at = :u"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberL{Value: res},
			":u": &types.AttributeValueMemberS{Value: updateTime},
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

func AddAgentToCustomers(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	customerIDs := req.MultiValueQueryStringParameters["id"]
	var data struct {
		AgentID []int `json:"agents_id"`
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

	for _, v := range customerIDs {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: v},
			},
			UpdateExpression: aws.String("SET agents_id = list_append(agents_id, :t)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":t": &types.AttributeValueMemberL{Value: res},
			},
			ReturnValues: types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToAddAgentToMany, CustomerID = ", v, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToAddAgentToMany, CustomerID = " + v + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func EditCustomersAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		OldAgent string `json:"old_agent"`
		NewAgent string `json:"new_agent"`
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
			":g": &types.AttributeValueMemberN{Value: data.OldAgent},
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
	oldAgentInt, _ := strconv.Atoi(data.OldAgent)

	for _, v := range customers {

		//set index for update tag
		var agentIndex string
		for i, old := range v.AgentsID {
			if old == oldAgentInt {
				agentIndex = strconv.Itoa(i)
				break
			}
		}

		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String("set agents_id[" + agentIndex + "].field = :ng"),
			ReturnValues:     types.ReturnValueUpdatedNew,
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":ng": &types.AttributeValueMemberN{Value: data.NewAgent},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateAgentToMany, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateAgentToMany, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

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

	var tagStr string

	for k, v := range data.Agents {
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

		if k == 0 {
			tagStr = "REMOVE agents_id[" + strconv.Itoa(k) + "]"
		} else {
			tagStr += ", agents_id[" + strconv.Itoa(k) + "]"
		}

	}

	updateTime := time.Now().Format(os.Getenv("TIMEFORMAT"))

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.CustomerId)},
		},
		UpdateExpression:    aws.String(tagStr + ", Set update_at = :u"),
		ReturnValues:        types.ReturnValueUpdatedNew,
		ConditionExpression: aws.String("attribute_exists(customer_id)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":u": &types.AttributeValueMemberS{Value: updateTime},
		},
	})
	if err != nil {
		fmt.Println("FailedToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}

func DeleteCustomersAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		DeleteAgent string `json:"delete_agent"`
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
			":g": &types.AttributeValueMemberN{Value: data.DeleteAgent},
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
	deleteAgentInt, _ := strconv.Atoi(data.DeleteAgent)

	for _, v := range customers {

		//set update index
		var agentIndex string
		for i, old := range v.TagsID {
			if old == deleteAgentInt {
				agentIndex = strconv.Itoa(i)
				break
			}
		}

		removeStr := "remove agents_id[" + agentIndex + "]"
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"customer_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.CustomerID)},
			},
			UpdateExpression: aws.String(removeStr),
		})
		if err != nil {
			fmt.Println("FailedToRemoveAgent, CustomerID = ", v.CustomerID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToRemoveAgent, CustomerID = " + strconv.Itoa(v.CustomerID) + ", " + err.Error())}), nil
		}
	}

	return ApiResponse(http.StatusOK, nil), nil

}
