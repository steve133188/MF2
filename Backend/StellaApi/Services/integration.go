package Services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mf-stella-api/Model"
	"net/http"
	"net/url"
	"strings"
)

//####################
func GetIntegrations(w http.ResponseWriter, req *http.Request) {
	targetUrl := "https://api.stella.sanuker.com/vnext/rest/integrations"

	urlVal := url.Values{}
	urlVal.Set("access_token", Model.Token)

	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()
	fmt.Println(urlStr)

	newResp, err := http.Post(urlStr, "application/json", nil)
	if err != nil {
		HandleError(w, err)
	}
	defer newResp.Body.Close()
	fmt.Println(newResp)
	result, err := ioutil.ReadAll(newResp.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}
	HandleResult(w, result)
}

func PutIntegrations(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://api.stella.sanuker.com/vnext/rest/integration"

	urlVal := url.Values{}
	urlVal.Set("access_token", Model.Token)

	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	parseData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	newReq, err := http.NewRequest("PUT", urlStr, bytes.NewBuffer(parseData))
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

func DeleteIntegrations(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://api.stella.sanuker.com/vnext/rest/integration/"

	id := strings.TrimPrefix(req.URL.Path, "/get_integration/")
	urlVal := url.Values{}
	urlVal.Set("access_token", Model.Token)

	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + id + "?" + urlVal.Encode()
	fmt.Println(urlStr)

	newReq, err := http.NewRequest("DELETE", urlStr, nil)
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
