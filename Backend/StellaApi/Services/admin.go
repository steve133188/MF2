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

func GetAdmins(w http.ResponseWriter, req *http.Request) {

	targetUrl := "https://api.stella.sanuker.com/vnext/rest/admins/get"

	urlVal := url.Values{}

	for k, v := range req.URL.Query() {
		urlVal.Set(k, v[0])
	}
	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	data := new(Model.GetAdmins)

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to Marshal json to model")
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}
	fmt.Println(bytes.NewBuffer(parseData))
	newResp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(parseData))
	if err != nil {
		HandleError(w, err)
	}
	defer newResp.Body.Close()

	result, err := ioutil.ReadAll(newResp.Body)
	HandleResult(w, result)

}
