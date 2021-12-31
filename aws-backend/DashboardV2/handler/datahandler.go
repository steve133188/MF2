package handler

import (
	"context"
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

//Contact Info ************************************************************************************************

func GetAllContact(userID int, customers []model.Customer) int {
	var count int
	for _, v := range customers {
		for _, vAgent := range v.AgentsID {
			if vAgent == userID {
				count++
			}
		}
	}
	return count
}

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

func GetActiveContacts(userId int, customers []model.Customer, messages []model.Message) int {
	count := 0
	for _, v := range messages {
		for _, k := range customers {
			if k.CustomerID == v.RoomID && strings.Contains(v.Sender, strconv.Itoa(userId)) {
				count++
				break
			}
		}
	}
	return count
}

//Message Info ************************************************************************************************

func GetRespTime(userID int, messages []model.Message) (model.RespTime, error) {
	var respTime model.RespTime
	var respTimeArr []int64

	//get All response time
	i := 0
	for i < len(messages) {
		if strings.Contains(messages[i].Receiver, strconv.Itoa(userID)) {
			for j := i; j < len(messages); j++ {
				if messages[j].Sender == messages[i].Receiver {
					jtime, err := strconv.ParseInt(messages[j].TimeStamp, 10, 64)
					if err != nil {
						return model.RespTime{}, err
					}
					itime, err := strconv.ParseInt(messages[i].TimeStamp, 10, 64)
					if err != nil {
						return model.RespTime{}, err
					}
					respTimeArr = append(respTimeArr, jtime-itime)
				}
			}
		}
		i++
	}

	if len(respTimeArr) == 0 {
		return respTime, nil
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
	return respTime, nil
}

func GetTotalMsgSent(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if strings.Contains(v.Sender, strconv.Itoa(userID)) {
			count++
		}
	}
	return count
}

func GetTotalMsgRev(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if strings.Contains(v.Receiver, strconv.Itoa(userID)) {
			count++
		}
	}
	return count
}

func GetCommunicationNumber(userID int, messages []model.Message) int {
	var count int
	var countList []int
	for i, v := range messages {
		if strings.Contains(v.Receiver, strconv.Itoa(userID)) || strings.Contains(v.Sender, strconv.Itoa(userID)) {
			if len(countList) == 0 || intInSlice(messages[i].RoomID, countList) {
				countList = append(countList, messages[i].RoomID)
				count++
			}
		}
	}
	return count
}

//Tool functions ************************************************************************************************

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
