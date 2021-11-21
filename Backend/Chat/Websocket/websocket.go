package Websocket

import (
	"errors"
	"fmt"
	"log"
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
	// r.ParseForm()
	// roomID := r.Form["room_id"][0]
	// uid := r.Form["user_id"][0]

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}
	// defer ws.Close()

	// SetWsConn(roomID, uid, ws)

	client := NewClient(ws, ws.RemoteAddr().String())

	//should be placed in main.go??
	// go func() {
	// 	WhatsappRoute := mux.NewRouter()
	// 	WhatsappRoute.HandleFunc("/whatsapp", client.HandleWhatsapp)
	// 	err = http.ListenAndServe(":3013", WhatsappRoute)
	// 	if err != nil {
	// 		log.Fatal("3013", err)
	// 	}
	// }()

	go client.read()
	go client.write()

}

func SetWsConn(room string, user string, conn *websocket.Conn) {
	if Rooms[room] == nil {
		con := make(map[string][]*websocket.Conn)
		con[user] = append(con[user], conn)
		Rooms[room] = con
	} else {
		Rooms[room][user] = append(Rooms[room][user], conn)
	}
	for k1, v1 := range Rooms {
		fmt.Println(k1)
		fmt.Println(v1)

		for k, v := range v1 {
			fmt.Println(k)
			fmt.Println(v)
		}

	}
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
	for k, v := range Rooms[room][user] {
		if v == conn {
			Rooms[room][user] = append(Rooms[room][user][:k], Rooms[room][user][k+1])
			fmt.Println(Rooms[room][user])
		}
	}
}
