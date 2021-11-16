package Websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}
	// defer ws.Close()

	client := NewClient(ws, ws.RemoteAddr().String())

	//should be placed in main.go??
	go func() {
		WhatsappRoute := mux.NewRouter()
		WhatsappRoute.HandleFunc("/whatsapp", client.HandleWhatsapp)
		err = http.ListenAndServe(":3013", WhatsappRoute)
		if err != nil {
			log.Fatal("3013", err)
		}
	}()

	go client.read()
	go client.write()

}
