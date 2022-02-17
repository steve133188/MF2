package handler

import (
	"aws-lambda-standard-replies/model"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetStdReplyByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		fmt.Println("failed to getItem, ID = ", id)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to getItem, id = " + id)}), nil
	}
	if out.Item == nil {
		fmt.Println("ItemNotFound ID = ", id)
		return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ItemNotFound id = " + id)}), nil
	}

	stdReply := new(model.StandardReplies)

	err = attributevalue.UnmarshalMap(out.Item, &stdReply)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, ID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, ID =  " + id + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, stdReply), nil
}

func GetAllStdReplies(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	stdReplies := make([]model.StandardReplies, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName: aws.String(table),
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScan, " + err.Error())}), nil
		}

		pstdReplies := make([]model.StandardReplies, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pstdReplies)
		if err != nil {
			fmt.Println("FailedToUnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMaps, " + err.Error())}), nil
		}

		stdReplies = append(stdReplies, pstdReplies...)
	}

	return ApiResponse(http.StatusOK, stdReplies), nil
}

func GetStdRepliesByChannel(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	channel := req.PathParameters["channel"]

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		FilterExpression: aws.String("contains(#ch, :ch)"),
		ExpressionAttributeNames: map[string]string{
			"#ch": "channels",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ch": &types.AttributeValueMemberS{Value: channel},
		},
	})

	channels := make([]model.StandardReplies, 0)

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("failed to scan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to scan, " + err.Error())}), nil
		}

		pChannels := make([]model.StandardReplies, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pChannels)
		if err != nil {
			fmt.Println("failed to UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to UnmarshalListOfMaps, " + err.Error())}), nil
		}
		channels = append(channels, pChannels...)
	}

	return ApiResponse(http.StatusOK, channels), nil
}

func GetAllRepliesName(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	var data struct {
		Name []string `json:""`
	}

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:            aws.String(table),
		ProjectionExpression: aws.String("#name"),
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
		},
	})

	for p.HasMorePages() {
		pout, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("failed to scan, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to scan, " + err.Error())}), nil
		}

		pReplies := make([]model.StandardReplies, 0)
		err = attributevalue.UnmarshalListOfMaps(pout.Items, &pReplies)
		if err != nil {
			fmt.Println("failed to UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("failed to UnmarshalListOfMaps, " + err.Error())}), nil
		}

		for _, v := range pReplies {
			data.Name = append(data.Name, v.Name)
		}
	}

	return ApiResponse(http.StatusOK, data), nil
}
