package handler

import (
	"mf2-aws-dashboard/model"
	"strconv"
	"strings"
)

func GetAllContact(userID int, customers []model.Customer) int {
	var count int
	for _, v := range customers {
		for _, vAgent := range v.AgentsID {
			if vAgent == userID {
				count++
				break
			}
		}
	}
	return count
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

func GetRespTime(userID int, messages []model.Message) model.RespTime {
	var respTime model.RespTime
	var respTimeArr []int64

	//get All response time
	i := 0
	for i < len(messages) {
		if strings.Contains(messages[i].Receiver, strconv.Itoa(userID)) {
			for j := i; j < len(messages); j++ {
				if strings.Contains(messages[j].Sender, messages[i].Receiver) {
					respTimeArr = append(respTimeArr, messages[j].TimeStamp-messages[i].TimeStamp)
				}
			}
		}
		i++
	}

	if len(respTimeArr) == 0 {
		return respTime
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
	respTime.First = int(respTimeArr[0])
	respTime.Longest = int(max)
	respTime.Average = int(sum) / len(respTimeArr)
	return respTime
}

func GetCommunicationNumber(userID int, messages []model.Message) int {
	var count int
	var countList []int
	for i, v := range messages {
		if strings.Contains(v.Sender, strconv.Itoa(userID)) || strings.Contains(v.Receiver, strconv.Itoa(userID)) {
			if len(countList) == 0 || !intInSlice(messages[i].RoomID, countList) {
				countList = append(countList, messages[i].RoomID)
				count++
			}
		}
	}
	return count
}

func GetTags(userID int, customers []model.Customer, tags []model.Tag) map[string]int {
	var tagIDs map[int]int

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
	var tagNo map[string]int
	for _, v := range tags {
		if _, found := tagIDs[v.TagID]; found {
			tagNo[v.TagName] = tagIDs[v.TagID]
		}
	}
	return tagNo
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
