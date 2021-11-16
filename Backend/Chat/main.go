package main

import (
	"log"
	"mf-chat-services/DB"
	"mf-chat-services/Websocket"
	"net/http"
	// "mf-chat-services/Util"
	// jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	DB.MongoConnect()

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
	// }))

	http.HandleFunc("/websocket", Websocket.HandleConnections)

	err := http.ListenAndServe(":3003", nil)
	if err != nil {
		log.Fatal(err)
	}

}
