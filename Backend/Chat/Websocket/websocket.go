package Websocket

import (
	"errors"
	"log"
	"mf-chat-services/Services"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Rooms = make(map[string]map[string][]*websocket.Conn)

func HandleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}
	client := NewClient(ws, ws.RemoteAddr().String())

	go client.read()
	go client.write()

}

func SetWsConn(room string, user string, conn *websocket.Conn) {
	if Rooms[room] == nil {
		con := make(map[string][]*websocket.Conn)
		con[user] = append(con[user], conn)
		Rooms[room] = con
		err := Services.CreateChatRoom(room, user)
		if err != nil {
			log.Println(err)
		}
	} else {
		Rooms[room][user] = append(Rooms[room][user], conn)
	}
	log.Println(Rooms)
}

func GetWsConn(room string, user string, conn *websocket.Conn) ([]*websocket.Conn, error) {
	if room == "" {
		log.Println("No this Room")
		return nil, errors.New("no room id")
	}
	for roomKey, userKey := range Rooms {
		if roomKey != room {
			continue
		} else {
			return userKey[user], nil
		}
	}
	log.Println("No this user")
	return nil, errors.New("no user")
}

func DelWsConn(room string, user string, conn *websocket.Conn) {
	log.Println("delete conn")

	for k, v := range Rooms[room][user] {
		if v == conn {
			result := deleteConnArrayItem(Rooms[room][user], k)
			Rooms[room][user] = result
			log.Println("delete      ", Rooms)
			break
		}
	}
	if len(Rooms[room][user]) == 0 {
		delete(Rooms, room)
		err := Services.DeleteChatRoom(room, user)
		if err != nil {
			log.Println(err)
		}
	}
}

func deleteConnArrayItem(conns []*websocket.Conn, k int) []*websocket.Conn {
	conns[k] = conns[len(conns)-1]
	conns = conns[:len(conns)-1]
	return conns
}
