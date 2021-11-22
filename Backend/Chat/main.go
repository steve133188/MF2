package main

import (
	"log"
	"mf-chat-services/DB"
	"mf-chat-services/Websocket"
	"net/http"

	"github.com/gorilla/mux"
	// "mf-chat-services/Util"
	// jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	DB.MongoConnect()

	websocket := mux.NewRouter()

	websocket.HandleFunc("/websocket", Websocket.HandleConnections)

	websocket.HandleFunc("/whatsapp", Websocket.HandleWhatsapp)

	err := http.ListenAndServe(":3003", websocket)
	if err != nil {
		log.Fatal(err)
	}

}
