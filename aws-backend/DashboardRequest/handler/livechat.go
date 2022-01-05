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
	"sort"
)

func GetLivechatDashboard(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
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

	dashboard := make([]model.Dashboard, 0)
	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}
		pItems := make([]model.Dashboard, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		dashboard = append(dashboard, pItems...)
	}

	sort.Slice(dashboard, func(i, j int) bool {
		return dashboard[i].TimeStamp < dashboard[j].TimeStamp
	})

	livechat := new(model.Livechat)
	dashLen := len(dashboard)
	livechat.AllContacts = make([]int, dashLen)
	livechat.ActiveContacts = make([]int, dashLen)
	livechat.TotalMsgSent = make([]int, dashLen)
	livechat.TotalMsgRev = make([]int, dashLen)
	livechat.AvgRespTime = make([]int64, dashLen)
	livechat.NewAddedContacts = make([]int, dashLen)
	livechat.CommunicationHours = make([]int, dashLen)
	livechat.NewAddedContacts = make([]int, dashLen)

	livechat.Tags = dashboard[dashLen-1].Tags
	for _, v := range dashboard[dashLen-1].Channel {
		newChannel := new(model.ChannelContact)
		newChannel.ChannelName = v.ChannelName
		newChannel.ChannelTotalContact = v.AllContacts
		livechat.ChannelContact = append(livechat.ChannelContact, *newChannel)
	}

	for i, v := range dashboard {

		for len(dashboard) <= 31 {
			livechat.AllContacts[i] = v.AllContacts
			livechat.ActiveContacts[i] = v.TotalActiveContacts

			livechat.TotalMsgSent[i] = v.TotalMsgSent
			livechat.TotalMsgRev[i] = v.TotalMsgRev

			livechat.AvgRespTime[i] = v.AvgTotalRespTime
			livechat.AvgTotalRespTime += v.AvgTotalRespTime
			livechat.AvgTotalFirstRespTime += v.AvgTotalFirstRespTime
			livechat.AvgLongestRespTime += v.AvgLongestRespTime

			livechat.NewAddedContacts[i] = v.NewAddedContacts
			for j, k := range v.CommunicationHours {
				livechat.CommunicationHours[j] = k
			}

			livechat.Yaxis[i] = v.TimeStamp
		}
	}

	livechat.AvgTotalRespTime = livechat.AvgTotalRespTime / int64(dashLen)
	livechat.AvgTotalFirstRespTime = livechat.AvgTotalFirstRespTime / int64(dashLen)
	livechat.AvgLongestRespTime = livechat.AvgLongestRespTime / int64(dashLen)

	return ApiResponse(http.StatusOK, livechat), nil
}

func GetLivechatDashboardByChannel(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	//fmt.Println("req.Parameters = ", req.PathParameters)

	startDate := req.QueryStringParameters["start"]
	endDate := req.QueryStringParameters["end"]
	channel := req.QueryStringParameters["channel"]
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

	dashboard := make([]model.Dashboard, 0)
	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("ErrorInNextPage, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorInNextPage")}), nil
		}
		pItems := make([]model.Dashboard, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pItems)
		if err != nil {
			fmt.Println("UnmarshalListOfMaps, ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorUnmarshalListOfMapsInNextPage")}), nil
		}

		dashboard = append(dashboard, pItems...)
	}

	sort.Slice(dashboard, func(i, j int) bool {
		return dashboard[i].TimeStamp < dashboard[j].TimeStamp
	})

	var channelIndex int
	for i, v := range dashboard[0].Channel {
		if v.ChannelName == channel {
			channelIndex = i
			break
		}
	}
	if channelIndex == 0 {
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ChannelNotFound")}), nil
	}

	channelLivechat := new(model.Livechat)
	dashLen := len(dashboard)
	channelLivechat.AllContacts = make([]int, dashLen)
	channelLivechat.ActiveContacts = make([]int, dashLen)
	channelLivechat.TotalMsgSent = make([]int, dashLen)
	channelLivechat.TotalMsgRev = make([]int, dashLen)
	channelLivechat.AvgRespTime = make([]int64, dashLen)
	channelLivechat.NewAddedContacts = make([]int, dashLen)
	channelLivechat.CommunicationHours = make([]int, dashLen)
	channelLivechat.NewAddedContacts = make([]int, dashLen)

	channelLivechat.Tags = dashboard[len(dashboard)-1].Tags
	for _, v := range dashboard[len(dashboard)-1].Channel {
		newChannel := new(model.ChannelContact)
		newChannel.ChannelName = v.ChannelName
		newChannel.ChannelTotalContact = v.AllContacts
		channelLivechat.ChannelContact = append(channelLivechat.ChannelContact, *newChannel)
	}

	for i, v := range dashboard {
		for len(dashboard) <= 31 {
			channelLivechat.AllContacts[i] = v.Channel[channelIndex].AllContacts
			channelLivechat.ActiveContacts[i] = v.Channel[channelIndex].TotalActiveContacts

			channelLivechat.TotalMsgSent[i] = v.Channel[channelIndex].TotalMsgSent
			channelLivechat.TotalMsgRev[i] = v.Channel[channelIndex].TotalMsgRev

			channelLivechat.AvgRespTime[i] = v.Channel[channelIndex].AvgTotalRespTime
			channelLivechat.AvgTotalRespTime += v.Channel[channelIndex].AvgTotalRespTime
			channelLivechat.AvgTotalFirstRespTime += v.Channel[channelIndex].AvgTotalFirstRespTime
			channelLivechat.AvgLongestRespTime += v.Channel[channelIndex].AvgLongestRespTime

			channelLivechat.NewAddedContacts[i] = v.Channel[channelIndex].NewAddedContacts
			for j, k := range v.Channel[channelIndex].CommunicationHours {
				channelLivechat.CommunicationHours[j] = k
			}

			channelLivechat.Yaxis[i] = v.TimeStamp
		}
	}
	channelLivechat.AvgTotalRespTime = channelLivechat.AvgTotalRespTime / int64(dashLen)
	channelLivechat.AvgTotalFirstRespTime = channelLivechat.AvgTotalFirstRespTime / int64(dashLen)
	channelLivechat.AvgLongestRespTime = channelLivechat.AvgLongestRespTime / int64(dashLen)

	return ApiResponse(http.StatusOK, channelLivechat), nil
}
