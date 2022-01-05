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

func GetAgentDashboard(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
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

	agents := new(model.Agents)
	last := len(dashboard) - 1
	agents.DataCollected = last + 1

	agents.TotalMsgSent = dashboard[last].TotalMsgSent
	agents.TotalMsgRev = dashboard[last].TotalMsgRev

	agents.TotalActiveContacts = dashboard[last].TotalActiveContacts
	agents.TotalAssignedContacts = dashboard[last].TotalAssignedContacts
	agents.TotalDeliveredContacts = dashboard[last].TotalDeliveredContacts
	agents.TotalUnhandledContact = dashboard[last].TotalUnhandledContact

	agents.AvgTotalRespTime = dashboard[last].AvgTotalRespTime
	agents.AvgTotalFirstRespTime = dashboard[last].AvgTotalFirstRespTime
	agents.AvgLongestRespTime = dashboard[last].AvgLongestRespTime

	agents.Agent = make([]model.Agent, 0)
	agents.AgentsNo = make([]int, last+1)
	agents.Connected = make([]int, last+1)
	agents.Disconnected = make([]int, last+1)
	agents.AllContacts = make([]int, last+1)
	agents.NewAddedContacts = make([]int, last+1)

	for i, v := range dashboard {
		agents.AllContacts[i] = dashboard[i].AllContacts
		agents.NewAddedContacts[i] = dashboard[i].NewAddedContacts
		agents.AgentsNo[i] = len(dashboard[i].User)

		if last+1 <= 31 {
			for j, k := range v.User {
				if i == 0 {
					newAgent := new(model.Agent)

					if k.UserStatus == "Connected" {
						agents.Connected[i]++
					} else if k.UserStatus == "Disconnected" {
						agents.Disconnected[i]++

					}

					newAgent.AssignedContacts = k.AssignedContacts
					newAgent.ActiveContacts = k.ActiveContacts
					newAgent.DeliveredContacts = k.DeliveredContacts
					newAgent.UnhandledContact = k.UnhandledContact

					newAgent.NewAddedContacts = k.NewAddedContacts
					newAgent.AllContacts = k.AllContacts

					newAgent.AvgRespTime = k.AvgRespTime
					newAgent.FirstRespTime = k.FirstRespTime
					newAgent.LongestRespTime = k.LongestRespTime

					newAgent.MsgSent = k.MsgSent
					newAgent.MsgRev = k.MsgRev

					agents.Agent = append(agents.Agent, *newAgent)
				} else {

					agents.Agent[j].NewAddedContacts += k.NewAddedContacts
					agents.Agent[j].AllContacts += k.AllContacts

					agents.Agent[i].AssignedContacts += k.AssignedContacts
					agents.Agent[i].ActiveContacts += k.ActiveContacts
					agents.Agent[i].DeliveredContacts += k.DeliveredContacts
					agents.Agent[i].UnhandledContact += k.UnhandledContact

					agents.Agent[j].AvgRespTime += k.AvgRespTime
					agents.Agent[j].FirstRespTime += k.FirstRespTime
					agents.Agent[j].LongestRespTime += k.LongestRespTime

					agents.Agent[j].MsgSent += k.MsgSent
					agents.Agent[j].MsgRev += k.MsgRev
				}
			}
		}
	}
	for i, _ := range agents.Agent {
		agents.Agent[i].UserID = dashboard[last].User[i].UserID
		agents.Agent[i].UserName = dashboard[last].User[i].UserName
		agents.Agent[i].UserRoleName = dashboard[last].User[i].UserRoleName
		agents.Agent[i].TeamID = dashboard[last].User[i].TeamID
		agents.Agent[i].UserStatus = dashboard[last].User[i].UserStatus

		agents.Agent[i].AvgRespTime = agents.Agent[i].AvgRespTime / int64(last+1)
		agents.Agent[i].FirstRespTime = agents.Agent[i].FirstRespTime / int64(last+1)
		agents.Agent[i].LongestRespTime = agents.Agent[i].LongestRespTime / int64(last+1)
	}

	return ApiResponse(http.StatusOK, agents), nil
}

func GetAgentDashboardByChannel(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
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

	agents := new(model.Agents)
	last := len(dashboard) - 1
	agents.DataCollected = last + 1

	agents.TotalMsgSent = dashboard[last].Channel[channelIndex].TotalMsgSent
	agents.TotalMsgRev = dashboard[last].Channel[channelIndex].TotalMsgRev

	agents.TotalActiveContacts = dashboard[last].Channel[channelIndex].TotalActiveContacts
	agents.TotalAssignedContacts = dashboard[last].Channel[channelIndex].TotalAssignedContacts
	agents.TotalDeliveredContacts = dashboard[last].Channel[channelIndex].TotalDeliveredContacts
	agents.TotalUnhandledContact = dashboard[last].Channel[channelIndex].TotalUnhandledContact

	agents.AvgTotalRespTime = dashboard[last].Channel[channelIndex].AvgTotalRespTime
	agents.AvgTotalFirstRespTime = dashboard[last].Channel[channelIndex].AvgTotalFirstRespTime
	agents.AvgLongestRespTime = dashboard[last].Channel[channelIndex].AvgLongestRespTime

	agents.Agent = make([]model.Agent, 0)
	agents.AgentsNo = make([]int, last+1)
	agents.Connected = make([]int, last+1)
	agents.Disconnected = make([]int, last+1)
	agents.AllContacts = make([]int, last+1)
	agents.NewAddedContacts = make([]int, last+1)

	for i, v := range dashboard {
		agents.AllContacts[i] = dashboard[i].Channel[channelIndex].AllContacts
		agents.NewAddedContacts[i] = dashboard[i].Channel[channelIndex].NewAddedContacts
		agents.AgentsNo[i] = len(dashboard[i].User)

		if last+1 <= 31 {
			for j, k := range v.Channel[channelIndex].User {
				if i == 0 {
					newAgent := new(model.Agent)

					if k.UserStatus == "Connected" {
						agents.Connected[i]++
					} else if k.UserStatus == "Disconnected" {
						agents.Disconnected[i]++

					}

					newAgent.AssignedContacts = k.AssignedContacts
					newAgent.ActiveContacts = k.ActiveContacts
					newAgent.DeliveredContacts = k.DeliveredContacts
					newAgent.UnhandledContact = k.UnhandledContact

					newAgent.NewAddedContacts = k.NewAddedContacts
					newAgent.AllContacts = k.AllContacts

					newAgent.AvgRespTime = k.AvgRespTime
					newAgent.FirstRespTime = k.FirstRespTime
					newAgent.LongestRespTime = k.LongestRespTime

					newAgent.MsgSent = k.MsgSent
					newAgent.MsgRev = k.MsgRev

					agents.Agent = append(agents.Agent, *newAgent)
				} else {
					agents.Agent[j].AssignedContacts += k.AssignedContacts
					agents.Agent[j].ActiveContacts += k.ActiveContacts
					agents.Agent[j].DeliveredContacts += k.DeliveredContacts
					agents.Agent[j].UnhandledContact += k.UnhandledContact

					agents.Agent[j].NewAddedContacts += k.NewAddedContacts
					agents.Agent[j].AllContacts += k.AllContacts

					agents.Agent[j].AvgRespTime += k.AvgRespTime
					agents.Agent[j].FirstRespTime += k.FirstRespTime
					agents.Agent[j].LongestRespTime += k.LongestRespTime

					agents.Agent[j].MsgSent += k.MsgSent
					agents.Agent[j].MsgRev += k.MsgRev
				}
			}
		}
	}

	for i, _ := range agents.Agent {
		agents.Agent[i].UserID = dashboard[last].User[i].UserID
		agents.Agent[i].UserName = dashboard[last].User[i].UserName
		agents.Agent[i].UserRoleName = dashboard[last].User[i].UserRoleName
		agents.Agent[i].TeamID = dashboard[last].User[i].TeamID
		agents.Agent[i].UserStatus = dashboard[last].User[i].UserStatus

		agents.Agent[i].AvgRespTime = agents.Agent[i].AvgRespTime / int64(last+1)
		agents.Agent[i].FirstRespTime = agents.Agent[i].FirstRespTime / int64(last+1)
		agents.Agent[i].LongestRespTime = agents.Agent[i].LongestRespTime / int64(last+1)
	}

	return ApiResponse(http.StatusOK, agents), nil
}
