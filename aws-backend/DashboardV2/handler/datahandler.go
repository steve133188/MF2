package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"mf2-lambda-dashboardv2/model"
	"os"
	"strconv"
	"strings"
)

//User Info ************************************************************************************************

func GetRoleName(dynaClient *dynamodb.Client, roleId int) (string, error) {
	role := new(model.Role)
	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("ROLETABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(roleId)},
		},
	})
	if err != nil {
		return "", err
	}
	err = attributevalue.UnmarshalMap(out.Item, &role)
	if err != nil {
		return "", err
	}
	return role.RoleName, nil
}

func GetTags(userID int, customers []model.Customer, tags []model.Tag) []model.Tags {
	tagIDs := make(map[int]int)
	//find out all the tags
	for _, v := range customers {
		if intInSlice(userID, v.AgentsID) {
			for _, vTag := range v.TagsID {
				if _, found := tagIDs[vTag]; found {
					tagIDs[vTag]++
				} else {
					tagIDs[vTag] = 1
				}
			}
		}
	}

	//set tag ID to tag name
	tagNo := make([]model.Tags, 0)
	for _, v := range tags {
		if _, found := tagIDs[v.TagID]; found {
			temp := new(model.Tags)
			temp.Name = v.TagName
			temp.No = tagIDs[v.TagID]
			tagNo = append(tagNo, *temp)
		}
	}
	return tagNo
}

func GetAllTags(user []model.User, customers []model.Customer, tags []model.Tag) []model.Tags {
	tagIDs := make(map[int]int)
	//find out all the tags
	for _, k := range user {
		for _, v := range customers {
			if intInSlice(k.UserID, v.AgentsID) {
				for _, vTag := range v.TagsID {
					if _, found := tagIDs[vTag]; found {
						tagIDs[vTag]++
					} else {
						tagIDs[vTag] = 1
					}
				}
			}
		}
	}

	//set tag ID to tag name
	tagNo := make([]model.Tags, 0)
	for _, v := range tags {
		if _, found := tagIDs[v.TagID]; found {
			temp := new(model.Tags)
			temp.Name = v.TagName
			temp.No = tagIDs[v.TagID]
			tagNo = append(tagNo, *temp)
		}
	}
	return tagNo
}

//Contact Info ************************************************************************************************

func GetAssignedContacts(userId int, customers []model.Customer) (int, []model.Customer) {
	count := 0
	data := make([]model.Customer, 0)
	for _, v := range customers {
		for _, j := range v.AgentsID {
			if j == userId {
				count++
				data = append(data, v)
				break
			}
		}
	}
	return count, data
}

func GetDeliveredContacts(userId int, customers []model.Customer, messages []model.Message) int {
	var contact []int
	count := 0
	for i, v := range messages {
		for _, k := range customers {
			sender, err := strconv.Atoi(messages[i].Sender)
			reciever, err := strconv.Atoi(messages[i].Recipient)
			if err != nil {
				return count
			}
			if k.CustomerID == reciever && strings.Contains(v.Sender, strconv.Itoa(userId)) {
				if !intInSlice(sender, contact) {
					contact = append(contact, sender)
					count++
				}
				if !intInSlice(k.CustomerID, contact) {
					contact = append(contact, k.CustomerID)
					count++
				}
				break
			}
		}
	}
	return count
}

//Message Info ************************************************************************************************

func GetRespTime(userID int, messages []model.Message) (model.RespTime, int, int, error) {
	var respTime model.RespTime
	var respTimeArr []int64
	var activeContactName []string
	var activeContactCount int
	var unhandledContactName []string
	var unhandledContactCount int

	//get All response time
	i := 0
	for i < len(messages) {
		if strings.Contains(messages[i].Recipient, strconv.Itoa(userID)) {
			for j := i; j < len(messages); j++ {
				if messages[j].Sender == messages[i].Recipient {
					jtime, err := strconv.ParseInt(messages[j].TimeStamp, 10, 64)
					if err != nil {
						return model.RespTime{}, activeContactCount, unhandledContactCount, err
					}
					itime, err := strconv.ParseInt(messages[i].TimeStamp, 10, 64)
					if err != nil {
						return model.RespTime{}, activeContactCount, unhandledContactCount, err
					}
					respTimeArr = append(respTimeArr, jtime-itime)

					if !strInSlice(messages[j].Sender, activeContactName) {
						activeContactName = append(activeContactName, messages[j].Sender)
						activeContactCount++
					}
					if !strInSlice(messages[i].Recipient, activeContactName) {
						activeContactName = append(activeContactName, messages[j].Sender)
						activeContactCount++
					}
					break
				}
				if j+1 == len(messages) {
					if !strInSlice(messages[j].Sender, unhandledContactName) {
						activeContactName = append(unhandledContactName, messages[j].Sender)
						unhandledContactCount++
					}
					if !strInSlice(messages[i].Recipient, unhandledContactName) {
						activeContactName = append(unhandledContactName, messages[j].Sender)
						unhandledContactCount++
					}
				}
			}
		}
		i++
	}

	if len(respTimeArr) == 0 {
		return respTime, 0, 0, nil
	}

	//calculate the sum and maximum
	var sum, max int64
	for i, v := range respTimeArr {
		sum += v
		if i == 0 || v > max {
			max = v
		}
	}

	// fill respTime
	respTime.First = respTimeArr[0]
	respTime.Longest = max
	respTime.Average = sum / int64(len(respTimeArr))
	return respTime, activeContactCount, unhandledContactCount, nil
}

func GetMsgSent(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if strings.Contains(v.Sender, strconv.Itoa(userID)) {
			count++
		}
	}
	return count
}

func GetMsgRev(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if strings.Contains(v.Recipient, strconv.Itoa(userID)) {
			count++
		}
	}
	return count
}

func GetTotalMsgSent(messages []model.Message) int {
	var count int
	for _, v := range messages {
		if v.FromMe {
			count++
		}
	}
	return count
}

//func GetCommunicationNumber(userID int, messages []model.Message) int {
//	var count int
//	var countList []int
//	for i, v := range messages {
//		if strings.Contains(v.Receiver, strconv.Itoa(userID)) || strings.Contains(v.Sender, strconv.Itoa(userID)) {
//			if len(countList) == 0 || intInSlice(messages[i].RoomID, countList) {
//				countList = append(countList, messages[i].RoomID)
//				count++
//			}
//		}
//	}
//	return count
//}

//Communication Time ************************************************************************************************

func GetCommunicationHour(start int64, end int64, noOfInterval int, message []model.Message) []int {
	timeInterval := (end - start) / int64(noOfInterval)
	noOfMsg := make([]int, noOfInterval)

	for _, v := range message {
		for i, _ := range noOfMsg {
			tempTime, err := strconv.ParseInt(v.TimeStamp, 10, 64)
			if err != nil {
				fmt.Println("Error type, ", err)
				return noOfMsg
			}
			if (tempTime >= (start + int64(i)*timeInterval)) && (tempTime <= (start + int64(i+1)*timeInterval)) {
				fmt.Println("Time found")
				noOfMsg[i]++
				break
			}
			fmt.Println("Compare:", tempTime, (start + int64(i)*timeInterval), (start + int64(i+1)*timeInterval))
		}
	}
	return noOfMsg
}

//Tool functions ************************************************************************************************

func intInSlice(a int, list []int) bool {
	if len(list) == 0 {
		return false
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func strInSlice(a string, list []string) bool {
	if len(list) == 0 {
		return false
	}

	for _, b := range list {
		if strings.Contains(b, a) {
			return true
		}
	}
	return false
}
