package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mf-stella-api/Model"
	"net/http"
	"net/url"
)

func SendResponse(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://bot.stellabot.com/sendResponses"

	urlVal := url.Values{}

	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	data := new(Model.BotRequestBody)

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to Marshal json to data")
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	newResp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(parseData))
	if err != nil {
		HandleError(w, err)
	}

	defer newResp.Body.Close()
	fmt.Println(newResp)

	result, err := ioutil.ReadAll(newResp.Body)
	HandleResult(w, result)

}

// func RedirectNode(w http.ResponseWriter, req *http.Request) {
// 	var targetUrl = "https://bot.stellabot.com/redirectMemberToNode"

// 	urlVal := url.Values{}

// 	urlVal.Set("access_token", Model.Token)
// 	u, _ := url.ParseRequestURI(targetUrl)
// 	urlStr := u.String() + "?" + urlVal.Encode()

// 	data := new(Model.BotRedirectNode)

// 	parseData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Failed to Marshal json to model")
// 	}

// 	parseData, err = ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		fmt.Println("Failed to read all")
// 	}

// 	newResp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(parseData))
// 	if err != nil {
// 		HandleError(w, err)
// 	}
// 	defer newResp.Body.Close()

// 	fmt.Println(newResp)

// 	result, err := ioutil.ReadAll(newResp.Body)
// 	HandleResult(w, result)

// }
