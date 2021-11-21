package Websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"mf-chat-services/Model"
	"net"

	"github.com/gorilla/websocket"
)

type Client struct {
	Socket  *websocket.Conn
	Write   chan []byte
	Read    chan []byte
	Address string
	UserID  string
	RoomID  string

	// Message interface{}
}

var Register = make(map[net.Addr]bool)

func NewClient(socket *websocket.Conn, address string) (client *Client) {
	client = &Client{
		Socket:  socket,
		Write:   make(chan []byte, 100),
		Read:    make(chan []byte, 100),
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
		close(c.Read)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Println("Failed to read client message: ", c, err)

			return
		}
		log.Println("Processing client message: ", string(message))
		c.Read <- message
		c.WebsocketProcessData()

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
		close(c.Write)
		DelWsConn(c.RoomID, c.UserID, c.Socket)
		fmt.Println("map   ", Register)
		log.Println("client write (defer)", c)
	}()

	for {
		message, ok := <-c.Write
		if !ok {
			log.Println("Nothing to Write: ", c, "ok", ok)

			return
		}
		c.Socket.WriteMessage(websocket.TextMessage, message)
		log.Println("Write message: ", string(message))
	}
}

func (c *Client) WebsocketProcessData() {
	message := <-c.Read
	if string(message) == "ping" {
		message = []byte("pong")
		c.Write <- message
	} else {
		c.Write <- []byte("message received")

		clientMsg := new(Model.ClientMsg)
		err := json.Unmarshal(message, &clientMsg)
		if err != nil {
			log.Println("process data     ", err)
			c.Write <- []byte(message)

		}
		recData := new(Model.ClientMsg)
		switch clientMsg.ChannelType {
		case "whatsapp":
			recData, err = c.HandleCliWhatsappMsg(clientMsg)
			if err != nil {
				log.Println("process data     ", err)
				c.Write <- []byte(err.Error())

			}
		default:
			c.Write <- []byte(message)

		}

		room := recData.ChatID
		user := recData.UserID

		//append conn to Rooms if necessary
		var found bool
		connSlice := Rooms[room][user]
		for _, v := range connSlice {
			if v == c.Socket {
				found = true
			}
		}
		if !found {
			SetWsConn(room, user, c.Socket)
		}
		SetWsConn(room, user, c.Socket)

		result, err := json.Marshal(recData)
		if err != nil {
			log.Println(err)
			c.Write <- []byte(err.Error())
		}
		c.Write <- []byte(result)
		c.Write <- []byte("message sent")

		// http.ListenAndServe(Util.GoDotEnvVariable("WHATSAPP_ADRESS"), nil)
	}

}

// func (c *Client) HandleWhatsapp(w http.ResponseWriter, r *http.Request) {
// 	dec := json.NewDecoder(r.Body)
// 	data := new(Model.Chat)
// 	err := dec.Decode(&data)
// 	if err != nil {
// 		log.Println("Error in HandleWhatsapp: ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	message, err := json.Marshal(data)
// 	if err != nil {
// 		log.Println("Error in HandleWhatsapp marshal: ", err)
// 	}

// 	c.WriteMsg(message)
// 	log.Println(bytes.NewBuffer(message))
// }

// func (c *Client) WriteMsg(message []byte) {
// 	c.Write <- message
// }
