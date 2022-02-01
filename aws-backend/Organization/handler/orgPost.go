package handler

import (
	"aws-lambda-org/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"aws-lambda-org/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AddOrgItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	org := new(model.Organization)

	err := json.Unmarshal([]byte(req.Body), &org)
	if err != nil {
		log.Printf("FailedToUnmarshalInputData: %s", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalInputData")}), nil
	}

	fmt.Println(org.Type)

	if org.Type != "division" && org.Type != "team" {
		log.Println("WrongInputOfType, type should be division or team")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("WrongInputOfType, type should be division or team")}), nil

	}

	org.OrgID = utils.IdGenerator()

	if len(org.ChildrenID) == 0 {
		org.ChildrenID = make([]int, 0)
	}

	if org.ParentID != 0 {
		var idStr []int
		idStr = append(idStr, org.OrgID)

		idList, err := attributevalue.MarshalList(idStr)
		if err != nil {
			fmt.Println("FailedToMarshalListOrgID, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalListOrgID, " + err.Error())}), nil
		}
		out, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(org.ParentID)},
			},
			UpdateExpression: aws.String("Set children_id = list_append(children_id, :id)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":id": &types.AttributeValueMemberL{Value: idList},
			},
			ConditionExpression: aws.String("attribute_exists(org_id)"),
			ReturnValues:        types.ReturnValueAllNew,
		})
		if err != nil {
			fmt.Println("FailedToUpdateParent, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateParent, " + err.Error())}), nil
		}
		fmt.Println("out.attributes = ", out.Attributes)
	}

	av, err := attributevalue.MarshalMap(&org)
	if err != nil {
		log.Printf("FailedToMarshalMap: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMap")}), nil
	}

	out, err := dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(#id)"),
		ExpressionAttributeNames: map[string]string{
			"#id": strconv.Itoa(org.OrgID),
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		log.Printf("ErrorToAddItem: %s", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorToAddItem")}), nil

	}

	attributevalue.UnmarshalMap(out.Attributes, &org)
	return ApiResponse(http.StatusCreated, org), nil
}
