package Websocket

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mf-chat-services/Model"
	"mf-chat-services/Util"
	"net/http"
)

func HandleWhatsapp(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	data := new(Model.Chat)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error in HandleWhatsapp: ", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error in HandleWhatsapp marshal: ", err)
	}

	// c.sendMsg(message)
	log.Println(bytes.NewBuffer(message))
}

func (c *Client) HandleCliWhatsappMsg(msg *Model.ClientMsg) (*Model.ClientMsg, error) {
	clt := http.Client{}
	result, err := json.Marshal(msg)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_1     ", err)
		return nil, err
	}
	req, err := http.NewRequest("POST", Util.GoDotEnvVariable("WHATSAPP_ADDRESS"), bytes.NewBuffer(result))
	if err != nil {
		log.Println("HandleCliWhatsappMsg_2     ", err)
		return nil, err

	}
	res, err := clt.Do(req)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_3     ", err)
		return nil, err

	}

	data := new(Model.ClientMsg)
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_4     ", err)
		return nil, err

	}
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_5     ", err)
		return nil, err

	}
	return data, nil

}
