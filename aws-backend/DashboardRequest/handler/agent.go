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

	agents.TotalMsgSent = make([]int, 2)
	agents.TotalMsgRev = make([]int, 2)

	agents.TotalActiveContacts = make([]int, 2)
	agents.TotalAssignedContacts = make([]int, 2)
	agents.TotalDeliveredContacts = make([]int, 2)
	agents.TotalUnhandledContact = make([]int, 2)

	agents.AvgTotalRespTime = make([]int64, 2)
	agents.AvgTotalFirstRespTime = make([]int64, 2)
	agents.AvgLongestRespTime = make([]int64, 2)

	if last > 0 {
		agents.TotalMsgSent[0] = dashboard[last].TotalMsgSent
		agents.TotalMsgRev[0] = dashboard[last].TotalMsgRev

		agents.TotalActiveContacts[0] = dashboard[last-1].TotalActiveContacts
		agents.TotalAssignedContacts[0] = dashboard[last-1].TotalAssignedContacts
		agents.TotalDeliveredContacts[0] = dashboard[last-1].TotalDeliveredContacts
		agents.TotalUnhandledContact[0] = dashboard[last-1].TotalUnhandledContact

		agents.AvgTotalRespTime[0] = dashboard[last-1].AvgTotalRespTime
		agents.AvgTotalFirstRespTime[0] = dashboard[last-1].AvgTotalFirstRespTime
		agents.AvgLongestRespTime[0] = dashboard[last-1].AvgLongestRespTime
	}

	agents.TotalMsgSent[1] = dashboard[last].TotalMsgSent
	agents.TotalMsgRev[1] = dashboard[last].TotalMsgRev

	agents.TotalActiveContacts[1] = dashboard[last].TotalActiveContacts
	agents.TotalAssignedContacts[1] = dashboard[last].TotalAssignedContacts
	agents.TotalDeliveredContacts[1] = dashboard[last].TotalDeliveredContacts
	agents.TotalUnhandledContact[1] = dashboard[last].TotalUnhandledContact

	agents.AvgTotalRespTime[1] = dashboard[last].AvgTotalRespTime
	agents.AvgTotalFirstRespTime[1] = dashboard[last].AvgTotalFirstRespTime
	agents.AvgLongestRespTime[1] = dashboard[last].AvgLongestRespTime

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
			for _, k := range v.User {
				if i == 0 || getAgentIndex(agents.Agent, k.UserID) == -1 {
					newAgent := new(model.Agent)

					if k.UserStatus == "Connected" {
						agents.Connected[i]++
					} else if k.UserStatus == "Disconnected" {
						agents.Disconnected[i]++

					}

					newAgent.UserID = k.UserID
					newAgent.UserName = k.UserName
					newAgent.UserRoleName = k.UserRoleName
					newAgent.TeamID = k.TeamID
					newAgent.UserStatus = k.UserStatus

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

					targetIndex := getAgentIndex(agents.Agent, k.UserID)

					agents.Agent[targetIndex].UserName = k.UserName
					agents.Agent[targetIndex].UserRoleName = k.UserRoleName
					agents.Agent[targetIndex].TeamID = k.TeamID
					agents.Agent[targetIndex].UserStatus = k.UserStatus

					agents.Agent[targetIndex].NewAddedContacts += k.NewAddedContacts
					agents.Agent[targetIndex].AllContacts += k.AllContacts

					agents.Agent[targetIndex].AssignedContacts += k.AssignedContacts
					agents.Agent[targetIndex].ActiveContacts += k.ActiveContacts
					agents.Agent[targetIndex].DeliveredContacts += k.DeliveredContacts
					agents.Agent[targetIndex].UnhandledContact += k.UnhandledContact

					agents.Agent[targetIndex].AvgRespTime += k.AvgRespTime
					agents.Agent[targetIndex].FirstRespTime += k.FirstRespTime
					agents.Agent[targetIndex].LongestRespTime += k.LongestRespTime

					agents.Agent[targetIndex].MsgSent += k.MsgSent
					agents.Agent[targetIndex].MsgRev += k.MsgRev
				}
			}
		}
	}
	agentLen := len(agents.Agent)
	agents.Chart.UserName = make([]string, agentLen)
	agents.Chart.UserID = make([]int, agentLen)
	agents.Chart.TeamID = make([]int, agentLen)
	agents.Chart.Active = make([]int, agentLen)
	agents.Chart.Delivered = make([]int, agentLen)
	agents.Chart.Unhandled = make([]int, agentLen)

	for i, _ := range agents.Agent {

		agents.Agent[i].AvgRespTime = agents.Agent[i].AvgRespTime / int64(last+1)
		agents.Agent[i].FirstRespTime = agents.Agent[i].FirstRespTime / int64(last+1)
		agents.Agent[i].LongestRespTime = agents.Agent[i].LongestRespTime / int64(last+1)

		agents.Chart.UserName[i] = agents.Agent[i].UserName
		agents.Chart.UserID[i] = agents.Agent[i].UserID
		agents.Chart.TeamID[i] = agents.Agent[i].TeamID
		agents.Chart.Active[i] = agents.Agent[i].ActiveContacts
		agents.Chart.Delivered[i] = agents.Agent[i].DeliveredContacts
		agents.Chart.Unhandled[i] = agents.Agent[i].UnhandledContact
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

	agents.TotalMsgSent = make([]int, 2)
	agents.TotalMsgRev = make([]int, 2)

	agents.TotalActiveContacts = make([]int, 2)
	agents.TotalAssignedContacts = make([]int, 2)
	agents.TotalDeliveredContacts = make([]int, 2)
	agents.TotalUnhandledContact = make([]int, 2)

	agents.AvgTotalRespTime = make([]int64, 2)
	agents.AvgTotalFirstRespTime = make([]int64, 2)
	agents.AvgLongestRespTime = make([]int64, 2)

	if last > 0 {
		agents.TotalMsgSent[0] = dashboard[last-1].TotalMsgSent
		agents.TotalMsgRev[0] = dashboard[last-1].TotalMsgRev

		agents.TotalActiveContacts[0] = dashboard[last-1].TotalActiveContacts
		agents.TotalAssignedContacts[0] = dashboard[last-1].TotalAssignedContacts
		agents.TotalDeliveredContacts[0] = dashboard[last-1].TotalDeliveredContacts
		agents.TotalUnhandledContact[0] = dashboard[last-1].TotalUnhandledContact

		agents.AvgTotalRespTime[0] = dashboard[last-1].AvgTotalRespTime
		agents.AvgTotalFirstRespTime[0] = dashboard[last-1].AvgTotalFirstRespTime
		agents.AvgLongestRespTime[0] = dashboard[last-1].AvgLongestRespTime
	}

	agents.TotalMsgSent[1] = dashboard[last].TotalMsgSent
	agents.TotalMsgRev[1] = dashboard[last].TotalMsgRev

	agents.TotalActiveContacts[1] = dashboard[last].TotalActiveContacts
	agents.TotalAssignedContacts[1] = dashboard[last].TotalAssignedContacts
	agents.TotalDeliveredContacts[1] = dashboard[last].TotalDeliveredContacts
	agents.TotalUnhandledContact[1] = dashboard[last].TotalUnhandledContact

	agents.AvgTotalRespTime[1] = dashboard[last].AvgTotalRespTime
	agents.AvgTotalFirstRespTime[1] = dashboard[last].AvgTotalFirstRespTime
	agents.AvgLongestRespTime[1] = dashboard[last].AvgLongestRespTime

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
			for _, k := range v.Channel[channelIndex].User {
				if i == 0 || getAgentIndex(agents.Agent, k.UserID) == -1 {
					newAgent := new(model.Agent)

					if k.UserStatus == "Connected" {
						agents.Connected[i]++
					} else if k.UserStatus == "Disconnected" {
						agents.Disconnected[i]++

					}

					newAgent.UserID = k.UserID
					newAgent.UserName = k.UserName
					newAgent.UserRoleName = k.UserRoleName
					newAgent.TeamID = k.TeamID
					newAgent.UserStatus = k.UserStatus

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
					targetIndex := getAgentIndex(agents.Agent, k.UserID)

					agents.Agent[targetIndex].UserName = k.UserName
					agents.Agent[targetIndex].UserRoleName = k.UserRoleName
					agents.Agent[targetIndex].TeamID = k.TeamID
					agents.Agent[targetIndex].UserStatus = k.UserStatus

					agents.Agent[targetIndex].AssignedContacts += k.AssignedContacts
					agents.Agent[targetIndex].ActiveContacts += k.ActiveContacts
					agents.Agent[targetIndex].DeliveredContacts += k.DeliveredContacts
					agents.Agent[targetIndex].UnhandledContact += k.UnhandledContact

					agents.Agent[targetIndex].NewAddedContacts += k.NewAddedContacts
					agents.Agent[targetIndex].AllContacts += k.AllContacts

					agents.Agent[targetIndex].AvgRespTime += k.AvgRespTime
					agents.Agent[targetIndex].FirstRespTime += k.FirstRespTime
					agents.Agent[targetIndex].LongestRespTime += k.LongestRespTime

					agents.Agent[targetIndex].MsgSent += k.MsgSent
					agents.Agent[targetIndex].MsgRev += k.MsgRev
				}
			}
		}
	}

	agentLen := len(agents.Agent)
	agents.Chart.UserName = make([]string, agentLen)
	agents.Chart.UserID = make([]int, agentLen)
	agents.Chart.TeamID = make([]int, agentLen)
	agents.Chart.Active = make([]int, agentLen)
	agents.Chart.Delivered = make([]int, agentLen)
	agents.Chart.Unhandled = make([]int, agentLen)

	for i, _ := range agents.Agent {

		agents.Agent[i].AvgRespTime = agents.Agent[i].AvgRespTime / int64(last+1)
		agents.Agent[i].FirstRespTime = agents.Agent[i].FirstRespTime / int64(last+1)
		agents.Agent[i].LongestRespTime = agents.Agent[i].LongestRespTime / int64(last+1)

		agents.Chart.UserName[i] = agents.Agent[i].UserName
		agents.Chart.UserID[i] = agents.Agent[i].UserID
		agents.Chart.TeamID[i] = agents.Agent[i].TeamID
		agents.Chart.Active[i] = agents.Agent[i].ActiveContacts
		agents.Chart.Delivered[i] = agents.Agent[i].DeliveredContacts
		agents.Chart.Unhandled[i] = agents.Agent[i].UnhandledContact
	}

	return ApiResponse(http.StatusOK, agents), nil
}

func getAgentIndex(agents []model.Agent, target int) int {
	for i, v := range agents {
		if v.UserID == target {
			return i
		}
	}
	return -1
}
