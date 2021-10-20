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

func UpdatedChannel(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://api.stella.sanuker.com/vnext/rest/channel/info"

	urlVal := url.Values{}
	urlVal.Set("access_token", Model.Token)
	fmt.Println(Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()
	fmt.Println(urlStr)

	data := new(Model.UpdateChannel)

	parseData, err := json.Marshal(data)
	if err != nil {
		HandleError(w, err)
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		HandleError(w, err)
	}

	newReq, err := http.NewRequest(http.MethodPatch, urlStr, bytes.NewBuffer(parseData))
	fmt.Println(bytes.NewBuffer(parseData))

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
