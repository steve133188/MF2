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

	startDate := req.QueryStringParameters["start"]
	endDate := req.QueryStringParameters["end"]
	fmt.Println(startDate, endDate)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("#n BETWEEN :sd AND :ed"),
		ExpressionAttributeNames: map[string]string{
			"#n": "timestamp",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sd": &types.AttributeValueMemberN{Value: startDate},
			":ed": &types.AttributeValueMemberN{Value: endDate},
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

	startDate := req.QueryStringParameters["start"]
	endDate := req.QueryStringParameters["end"]

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        &table,
		Limit:            aws.Int32(50),
		FilterExpression: aws.String("#n >= :st AND #n <= :et"),
		ExpressionAttributeNames: map[string]string{
			"#n": "timestamp",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":st": &types.AttributeValueMemberN{Value: startDate},
			":et": &types.AttributeValueMemberN{Value: endDate},
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
