package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mf-chat-services/Model"
	"net/http"

	"gopkg.in/mgo.v2"
)

type test struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/message", func(rw http.ResponseWriter, r *http.Request) {
		var data Model.Message
		json.NewDecoder(r.Body).Decode(&data)
		session, err := mgo.Dial("mongodb://localhost:27017")
		//
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

		// val, err := json.Marshal(data)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		var test1 test
		test1.Phone = "85292443663@c.us"
		test1.Message = "知咩料啦"

		val, err := json.Marshal(test1)
		if err != nil {
			fmt.Println(err)
		}

		// t, _ := json.Marshal(test)

		newReq, err := http.NewRequest(http.MethodPost, "https://mf-whatsapp-js.herokuapp.com/send-message", bytes.NewBuffer(val))
		if err != nil {
			fmt.Println(err)
		}
		clt := http.Client{}
		fmt.Println(bytes.NewBuffer(val))
		fmt.Println(clt.Do(newReq))

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
