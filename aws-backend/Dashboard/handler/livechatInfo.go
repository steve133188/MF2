package handler

import (
	"mf2-aws-dashboard/model"
	"strconv"
)

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

func GetTotalMsgSent(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if v.Sender == userID {
			count++
		}
	}
	return count
}

func GetTotalMsgRev(userID int, messages []model.Message) int {
	var count int
	for _, v := range messages {
		if v.Receiver == userID {
			count++
		}
	}
	return count
}

func GetRespTime(userID int, messages []model.Message) (model.RespTime, error) {
	var respTime model.RespTime
	var respTimeArr []int64

	//get All response time
	i := 0
	for i < len(messages) {
		if messages[i].Receiver == userID {
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
	respTime.First = int(respTimeArr[0])
	respTime.Longest = int(max)
	respTime.Average = int(sum) / len(respTimeArr)
	return respTime, nil
}

func GetCommunicationNumber(userID int, messages []model.Message) int {
	var count int
	var countList []int
	for i, v := range messages {
		if v.Sender == userID || v.Receiver == userID {
			if len(countList) == 0 || intInSlice(messages[i].RoomID, countList) {
				countList = append(countList, messages[i].RoomID)
				count++
			}
		}
	}
	return count
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

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
