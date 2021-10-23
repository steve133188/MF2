package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mf-bot-services/Model"
	"net/http"
	"sync"
)

var Mu = sync.Mutex{}
var Botlist = make(map[int]string)        // number : senderid
var Stages = make(map[string]interface{}) // senderid : []Stages
var Steps = make(map[string]string)       // senderid : location id in MongoDB
var Num = 0

func HandleWhatsappResponse(resp http.ResponseWriter, req *http.Request) {
	//send request body to Flow
	senderId := req.URL.Query().Get("sender_id")
	message := req.URL.Query().Get("message")

	fmt.Println(senderId, message)

	clt := http.Client{}
	data := new(Model.BotBody)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		fmt.Println("1 :", err)
	}

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("2 ", err)
	}

	newReq, err := http.NewRequest("POST", "http://localhost:3006", bytes.NewBuffer(parseData))
	if err != nil {
		fmt.Println(err)
	}

	_, err = clt.Do(newReq)
	if err != nil {
		fmt.Println(err)
	}

	//test
	// parData, err := json.Marshal(data.Stages)
	// if err != nil {
	// 	fmt.Println("3 ", err)
	// }

	// mapping
	// Mu.Lock()
	hasBot := false
	for _, v := range Botlist {
		if data.SenderId == v {
			hasBot = true
		}
	}
	if !hasBot {
		Botlist[Num] = senderId
		Stages[senderId] = data.Stages
		// Stages[data.SenderId] = parData
		Num++
	}

	// Mu.Unlock()
	sendResponse(senderId, message)

}

func sendResponse(target string, message string) {
	var todo = Stages[target]

	var stages []Model.Stages
	var stage Model.Stages

	// var actions []Model.Actions
	var action Model.Actions

	// get []stage data
	if sl, ok := todo.([]interface{}); ok {
		for _, v := range sl {
			mapData := v.(map[string]interface{})
			b, err := json.Marshal(mapData)
			if err != nil {
				fmt.Println("send response marshal error")
			}
			err = json.Unmarshal(b, &stage)
			if err != nil {
				fmt.Println("send response unmarshal error")
			}
			stages = append(stages, stage)
			// fmt.Println("/////////")
			// test, _ := json.Marshal(stage)
			// fmt.Println(k, bytes.NewBuffer(test))
		}
		// test, _ := json.Marshal(stages)
		// fmt.Println(bytes.NewBuffer(test))
	}

	//get get source id from stage
	Steps[target] = "12314"
	if Steps[target] == "" {
		if sl, ok := stages[0].Actions.([]interface{}); ok {
			for k, v := range sl {
				mapData := v.(map[string]interface{})
				b, err := json.Marshal(mapData)
				if err != nil {
					fmt.Println("send response marshal error")
				}
				err = json.Unmarshal(b, &action)
				if err != nil {
					fmt.Println("send response unmarshal error")
				}
				fmt.Println("/////////")
				test, _ := json.Marshal(action)
				fmt.Println(k, bytes.NewBuffer(test))
			}
			// test, _ := json.Marshal(stages)
			// fmt.Println(bytes.NewBuffer(test))
		}
		Steps[target] = action.Id
	} else {
		for i := 0; i < len(stages); i++ {
			if sl, ok := stages[i].Actions.([]interface{}); ok {
				for _, v := range sl {
					mapData := v.(map[string]interface{})
					b, err := json.Marshal(mapData)
					if err != nil {
						fmt.Println("send response marshal error")
					}
					err = json.Unmarshal(b, &action)
					if err != nil {
						fmt.Println("send response unmarshal error")
					}
					fmt.Println(i, " ", action.Id)
				}
				// test, _ := json.Marshal(stages)
				// fmt.Println(bytes.NewBuffer(test))
			}
		}
	}

	//testing
	// s := reflect.ValueOf(todo)
	// b := s.Interface().([]byte)

	// err := json.Unmarshal(b, &todos)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, v := range todos {
	// 	// fmt.Println(k, v.Actions)
	// 	// fmt.Println("//////////")
	// 	// fmt.Println(reflect.TypeOf(v.Actions))
	// 	val := reflect.ValueOf(v.Actions)
	// 	fmt.Println(val.Interface().([]byte))

	// }

	//send response
	//http.NewRequest("POST", "https://bot.stellabot.com/sendResponses")

}
