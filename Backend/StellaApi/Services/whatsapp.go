package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mf-stella-api/Model"
	"net/http"
	"net/url"
	"strings"
)

func GetWhatsappFile(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://api.stella.sanuker.com/vnext/file/whatsapp"

	id := strings.TrimPrefix(req.URL.Path, "/get_whatsapp_file/")

	urlVal := url.Values{}

	for k, v := range req.URL.Query() {
		urlVal.Set(k, v[0])
	}
	urlVal.Set("access_token", Model.Token)
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "/" + id + "?" + urlVal.Encode()

	fmt.Println(urlStr)

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

func GetWhatsappTemplate(w http.ResponseWriter, req *http.Request) {
	var targetUrl = "https://api.stella.sanuker.com/vnext/rest/message_templates/get"

	urlVal := url.Values{}

	for k, v := range req.URL.Query() {
		urlVal.Set(k, v[0])
	}
	urlVal.Set("access_token", "Model.Token")
	u, _ := url.ParseRequestURI(targetUrl)
	urlStr := u.String() + "?" + urlVal.Encode()

	data := new(Model.WhatsappTemplate)

	parseData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to Marshal json to model")
	}

	parseData, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read all")
	}

	fmt.Println(urlStr)
	newReq, err := http.NewRequest("GET", urlStr, bytes.NewBuffer(parseData))
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
