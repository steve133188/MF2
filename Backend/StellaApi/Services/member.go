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

func GetMembers(w http.ResponseWriter, req *http.Request) {
	targetUrl := "https://api.stella.sanuker.com/vnext/rest/members/get"

	urlVal := url.Values{}

	for k, v := range req.URL.Query() {
		urlVal.Set(k, v[0])
	}

	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	data := new(Model.GetMembers)

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to Marshal json to model")
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	resp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(parseData))
	if err != nil {
		HandleError(w, err)
	}
	// fmt.Println(resp)
	result, err := ioutil.ReadAll(resp.Body)

	HandleResult(w, result)
	defer resp.Body.Close()
}

func GetMember(w http.ResponseWriter, req *http.Request) {
	targetUrl := "https://api.stella.sanuker.com/vnext/rest/member"
	// clt := http.Client{}

	urlVal := url.Values{}

	for k, v := range req.URL.Query() {
		urlVal.Set(k, v[0])
	}

	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	newReq, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		HandleError(w, err)
	}
	clt := http.Client{}
	newResp, err := clt.Do(newReq)
	if err != nil {
		HandleError(w, err)
	}
	result, err := ioutil.ReadAll(newResp.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}
	HandleResult(w, result)

}

//##############
func ToggleLiveChat(w http.ResponseWriter, req *http.Request) {
	targetUrl := "https://api.stella.sanuker.com/vnext/rest/member/live_chat"

	data := new(Model.ToggleLiveChat)

	urlVal := url.Values{}
	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to marshal json to model")
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	fmt.Println(bytes.NewBuffer(parseData))

	newReq, err := http.NewRequest(http.MethodPatch, urlStr, bytes.NewBuffer(parseData))

	clt := http.Client{}
	newResp, err := clt.Do(newReq)
	if err != nil {
		HandleError(w, err)
	}
	result, err := ioutil.ReadAll(newResp.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	HandleResult(w, result)
}
