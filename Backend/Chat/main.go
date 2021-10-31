package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mf-chat-services/Model"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {

	http.HandleFunc("/message", func(rw http.ResponseWriter, r *http.Request) {
		var data Model.Message
		json.NewDecoder(r.Body).Decode(&data)
		session, err := mgo.Dial("mongodb://localhost:27017")
		// mongodb+srv://backend-api:248E176vFbD09zeB@backend-mongodb-9bf339a5.mongo.ondigitalocean.com/chat?authSource=admin&replicaSet=backend-mongodb&tls=true&tlsCAFile=./tlsCAFile/ca-certificate.cer
		if err != nil {
			panic(err)
		}
		defer session.Close()
		c := session.DB("chat").C("chats")
		err = c.Insert(data)
		if err != nil {
			fmt.Println(err)
		}
		rw.WriteHeader(http.StatusOK)

		val, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		newReq, err := http.NewRequest(http.MethodPost, "http://62e4-118-140-230-94.ngrok.io", bytes.NewBuffer(val))
		if err != nil {
			fmt.Println(err)
		}
		clt := http.Client{}
		clt.Do(newReq)

		// rw.Header().Set("Content-Type", "application/json")
		// rw.WriteHeader(http.StatusOK)
		// val, _ := json.Marshal(data)
		// rw.Write(val)
	})
	http.ListenAndServe(":3003", nil)
}

// app := fiber.New()

// app.Use(logger.New())
// app.Use(cors.New())

// DB.MongoConnect()

// app.Get("/test", func(c *fiber.Ctx) error {
// 	return c.JSON(fiber.Map{"code": 200, "message": "Hello, MF-Chats-Services"})
// })

// app.Post("/message", func(c *fiber.Ctx) error {
// 	collection := DB.MI.DBCol

// 	data := new(Model.Message)

// 	err := c.BodyParser(&data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	result, err := collection.InsertOne(c.Context(), data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// get the inserted data
// 	todo := &Model.Message{}
// 	query := bson.D{{Key: "_id", Value: result.InsertedID}}

// 	collection.FindOne(c.Context(), query).Decode(todo)

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 		"data": fiber.Map{
// 			"message": todo,
// 		},
// 	})
// })

// app.Get("/getmessage", func(c *fiber.Ctx) error {

// })

// // app.Use(jwtware.New(jwtware.Config{
// // 	SigningKey: []byte(Util.GoDotEnvVariable("Token_pwd")),
// // }))

// api := app.Group("/api")

// Routes.ChatRoute(api.Group("/messages"))

// app.Listen(":3003")
