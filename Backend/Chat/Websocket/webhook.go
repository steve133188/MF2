package Websocket

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mf-chat-services/Model"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleWhatsapp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("webhook read req body      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		data := new(Model.ClientMsg)

		err = json.Unmarshal(reqBody, &data)
		if err != nil {
			log.Println("webhook unmarshal      ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		roomId := data.ChatID
		userId := data.UserID

		val, ok := Rooms[roomId][userId]
		if !ok {
			log.Println("Error in finding conn in Rooms map")
			w.WriteHeader(http.StatusBadRequest)
			str, _ := json.Marshal("user not connected")
			w.Write(str)
		}

		msg, _ := json.Marshal(data)
		for _, v := range val {
			v.WriteMessage(websocket.TextMessage, msg)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (c *Client) HandleCliWhatsappMsg(msg *Model.ClientMsg) (*Model.ClientMsg, error) {
	log.Println("=======================HandleCliWhatsappMsg=======================")
	clt := http.Client{}
	result, err := json.Marshal(msg)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_1     ", err)
		return nil, err
	}
	log.Println("handler result msg      ", bytes.NewBuffer(result))
	msg.Url = msg.Url + "/send-message"
	req, err := http.NewRequest("POST", msg.Url, bytes.NewBuffer(result))
	if err != nil {
		log.Println("HandleCliWhatsappMsg_2     ", err)
		return nil, err

	}
	req.Header.Set("Content-Type", "application/json")
	res, err := clt.Do(req)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_3     ", err)
		return nil, err

	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_4     ", err)
		return nil, err

	}

	log.Println("resp body       ", bytes.NewBuffer(resBody))

	data := new(Model.ClientMsg)
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		log.Println("HandleCliWhatsappMsg_5     ", err)
		return nil, err

	}
	return data, nil

}
