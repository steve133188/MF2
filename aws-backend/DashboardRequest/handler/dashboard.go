package handler

import (
	"aws-lambda-dashboardrequest/model"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
)

func GetLiveChat(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	//fmt.Println("req.Parameters = ", req.PathParameters)

	startDate := req.PathParameters["start"]
	endData := req.PathParameters["end"]

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("PK = livechat AND TimeStamp >= :st AND TimeStamp <= :et"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":st": &types.AttributeValueMemberN{Value: startDate},
			":et": &types.AttributeValueMemberN{Value: endData},
		},
	})

	liveChats := make([]model.LiveChat, 0)
	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}
		pItems := make([]model.LiveChat, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		liveChats = append(liveChats, pItems...)
	}
	return ApiResponse(http.StatusOK, liveChats), nil
}

func GetAgent(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	//fmt.Println("req.Parameters = ", req.PathParameters)

	startDate := req.PathParameters["start"]
	endData := req.PathParameters["end"]

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("PK = agent AND TimeStamp >= :st AND TimeStamp <= :et"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":st": &types.AttributeValueMemberN{Value: startDate},
			":et": &types.AttributeValueMemberN{Value: endData},
		},
	})

	agents := make([]model.Agent, 0)
	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}
		pItems := make([]model.Agent, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		agents = append(agents, pItems...)
	}
	return ApiResponse(http.StatusOK, agents), nil
}
