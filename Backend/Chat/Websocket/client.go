package Websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"mf-chat-services/Model"
	"time"

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
		close(c.Write)

	}()

	c.Socket.SetReadDeadline(time.Now().Add(5 * time.Second))
	c.Socket.SetPongHandler(func(string) error { c.Socket.SetReadDeadline(time.Now().Add(5 * time.Second)); return nil })
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
	// c.Socket.SetWriteDeadline(time.Now().Add(6 * time.Second))
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
		c.Socket.Close()
		DelWsConn(c.RoomID, c.UserID, c.Socket)
		fmt.Println("map   ", Rooms)
		log.Println("close c.Write")
	}()

	for {
		select {
		case message, ok := <-c.Write:
			if !ok {
				log.Println("Nothing to Write: ", c, "ok", ok)

				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
			log.Println("Write message: ", string(message))
		case <-ticker.C:
			err := c.Socket.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}

	}
}

func (c *Client) WebsocketProcessData() error {
	message := <-c.Read

	if string(message) == "ping" {
		message = []byte("pong")
		c.Write <- message
		return nil
	}
	chat := new(Model.ClientMsg)
	err := json.Unmarshal(message, &chat)
	if err != nil {
		log.Println(err)
		return err
	}
	if chat.Topic == "send-message" {
		c.Write <- []byte("message received")

		clientMsg := new(Model.ClientMsg)
		err := json.Unmarshal(message, &clientMsg)
		if err != nil {
			log.Println("process data unmarshal    ", err)
			c.Write <- []byte(message)
			return err
		}
		recData := new(Model.ClientMsg)
		switch clientMsg.ChannelType {
		case "whatsapp":
			recData, err = c.HandleCliWhatsappMsg(clientMsg)
			if err != nil {
				log.Println("process data handlewts    ", err)
				c.Write <- []byte(err.Error())
				return err
			}
			c.Write <- []byte("message sent")

		default:
			c.Write <- []byte("message sent")

			c.Write <- []byte(message)
			return nil
		}

		result, err := json.Marshal(recData)
		if err != nil {
			log.Println(err)
			c.Write <- []byte(err.Error())
			return err
		}

		c.RoomID = recData.ChatID
		c.UserID = recData.UserID

		var found = false
		for _, conns := range Rooms[c.RoomID] {
			for _, v := range conns {
				if v == c.Socket {
					found = true
				} else {
					err = v.WriteMessage(websocket.TextMessage, result)
					if err != nil {
						log.Println("broadcast      ", err)
					}
				}

			}
		}
		if !found {
			SetWsConn(c.RoomID, c.UserID, c.Socket)

		}
		c.Write <- []byte(result)

	}
	return nil
}
