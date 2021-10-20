package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mf-bot-services/Model"
	"net/http"
	"reflect"
	"sync"
)

var Mu = sync.Mutex{}
var Botlist = make(map[int]string)
var Stages = make(map[string]interface{})
var Num = 0

func HandleWhatsappResponse(resp http.ResponseWriter, req *http.Request) {
	//send request body to Flow
	clt := http.Client{}
	data := new(Model.BotBody)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(bytes.NewBuffer(parseData))
	newReq, err := http.NewRequest("POST", "http://localhost:3006", bytes.NewBuffer(parseData))
	if err != nil {
		fmt.Println(err)
	}

	_, err = clt.Do(newReq)
	if err != nil {
		fmt.Println(err)
	}

	// mapping
	// Mu.Lock()
	hasBot := false
	for _, v := range Botlist {
		if data.SenderId == v {
			hasBot = true
		}
	}
	if hasBot == false {
		Botlist[Num] = data.SenderId
		Stages[data.SenderId] = data.Stages
		Num++
	}
	// Mu.Unlock()
	sendResponse(data.SenderId)

}

func sendResponse(target string) {

	data := new(Model.Stages)
	stages := Stages[target]
	// var result []struct{}

	// var todos []Model.Stages = make([]Model.Stages, 0)
	//interface -> []struct{}
	switch reflect.TypeOf(stages).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(stages)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
			step, err := json.Marshal(s.Index(i))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(step))

			err = json.Unmarshal(step, &data)
			if err != nil {
				fmt.Println(err)
			}
			// result = append(result, *bytes.NewBuffer(step))
		}
	}

	// todos = result
	//send response
	//http.NewRequest("POST", "https://bot.stellabot.com/sendResponses")

}
