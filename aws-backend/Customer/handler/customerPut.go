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
