package Services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mf-flowbuilder-services/DB"
	"mf-flowbuilder-services/Model"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var Num = 0
var Map = make(map[int]string)
var Mu = sync.Mutex{}

func Testing(c *fiber.Ctx) error {
	data := new(Model.BotBody)
	if err := c.BodyParser(&data); err != nil {
		log.Println(err)
		//return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		fmt.Println("Failed to parse Ctx body")
	}
	fmt.Println(data)
	return c.Status(fiber.StatusOK).SendString("Success")
}

func NewReqBody(c *fiber.Ctx) error {
	collection := DB.MI.DBCol
	data := new(Model.BotBody)

	if err := c.BodyParser(&data); err != nil {
		log.Println(err)
		//return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		fmt.Println("Failed to parse Ctx body")
	}

	// dataStr, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("Failed to marshal data")
	// }

	Mu.Lock()
	Map[Num] = string(data.BotName)
	Num++
	result, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	fmt.Println(result)
	Mu.Unlock()
	go SendMap()

	return c.Status(fiber.StatusOK).SendString("Success")
}

func SendMap() {
	Mu.Lock()
	defer Mu.Unlock()
	data, err := json.Marshal(Map)
	if err != nil {
		return
	}

	go func() {
		req, err := http.NewRequest("Post", "localhost:3005", bytes.NewBuffer(data))
		if err != nil {
			return
		}
		clt := http.Client{}
		resp, err := clt.Do(req)
		if err != nil {
			return
		}
		fmt.Println(resp.Body)
	}()

}
