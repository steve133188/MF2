package Websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mf-chat-services/Model"
	"mf-chat-services/Util"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Socket  *websocket.Conn
	Send    chan []byte
	Address string
	// Message interface{}
}

func NewClient(socket *websocket.Conn, address string) (client *Client) {
	client = &Client{
		Socket:  socket,
		Send:    make(chan []byte, 100),
		Address: address,
	}
	return
}

func (c *Client) read() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("read stop", string(debug.Stack()), r)

	// 	}
	// }()

	defer func() {
		log.Println("Close client channel", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Println("Failed to read client message: ", c, err)

			return
		}
		log.Println("Processing client message: ", string(message))
		ProcessData(c, message)

	}
}

func (c *Client) write() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("write stop", string(debug.Stack()), r)

	// 	}
	// }()

	defer func() {
		c.Socket.Close()
		log.Println("client write (defer)", c)
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			log.Println("Nothing to send: ", c, "ok", ok)

			return
		}
		c.Socket.WriteMessage(websocket.TextMessage, message)
		log.Println("Write message: ", string(message))
	}
}

func ProcessData(c *Client, message []byte) {
	if string(message) == "ping" {
		message = []byte("pong")
		c.sendMsg(message)
	} else {
		clt := http.Client{}
		req, err := http.NewRequest("POST", Util.GoDotEnvVariable("WHATSAPP_ADDRESS"), bytes.NewBuffer(message))
		if err != nil {
			log.Println("Error in webhook post", err)
		}
		res, err := clt.Do(req)
		if err != nil {
			log.Println("Error after webhook post", err)
		}
		fmt.Println(res.Body)
		c.sendMsg([]byte("message sent"))
		// http.ListenAndServe(Util.GoDotEnvVariable("WHATSAPP_ADRESS"), nil)
	}

}

func (c *Client) HandleWhatsapp(w http.ResponseWriter, r *http.Request) {
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

	c.sendMsg(message)
	log.Println(bytes.NewBuffer(message))
}

func (c *Client) sendMsg(message []byte) {
	c.Send <- message
}
