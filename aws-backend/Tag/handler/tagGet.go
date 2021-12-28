package handler

import (
	"aws-lambda-tag/model"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetTagByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]types.AttributeValue{
			"tag_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, TagID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem, TagID = " + id + ", " + err.Error())}), nil
	}
	if out.Item == nil {
		fmt.Println("TagNotExists, TagID = ", id, ", ", err)
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("TagNotExists, TagID = " + id + ", " + err.Error())}), nil
	}

	tag := new(model.Tag)
	err = attributevalue.UnmarshalMap(out.Item, &tag)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, TagID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, TagID = " + id + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, tag), nil
}

func GetAllTags(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	tags := make([]model.Tag, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Limit:     aws.Int32(50),
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan, " + err.Error())}), nil
		}

		pTags := make([]model.Tag, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pTags)
		if err != nil {
			fmt.Println("FailedToUnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMaps, " + err.Error())}), nil
		}

		tags = append(tags, pTags...)
	}

	return ApiResponse(http.StatusOK, tags), nil
}
