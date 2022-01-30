package handler

import (
	"aws-lambda-org/model"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutOrgItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	org := new(model.Organization)

	err := json.Unmarshal([]byte(req.Body), &org)
	if err != nil {
		log.Printf("FailedToUnmarshalInputData: %s", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("FailedToUnmarshalInputData")}), nil
	}

	_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(org.OrgID)},
		},
		UpdateExpression: aws.String("Set #n = :n"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{Value: org.Name},
		},
		ExpressionAttributeNames: map[string]string{
			"#n": "name",
		},
		ConditionExpression: aws.String("attribute_exists(org_id)"),
	})
	if err != nil {
		log.Println("ErrorToUpdateItem, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorToUpdateItem, " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, nil), nil
}
