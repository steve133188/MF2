package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"mf2-broadcast/model"
	"net/http"
	"net/url"
)

func SendBCTest(c *fiber.Ctx) error {
	bc := new(model.SendTemp)

	err := c.BodyParser(&bc)
	if err != nil {
		log.Println("body parser error,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	failedList := make([]model.CInfo, 0)

	for _, v := range bc.Customer {
		template := make(map[string]string)
		template["TemplateName"] = bc.TemplateName
		template["LanguageCode"] = bc.LanguageCode
		template["Phone"] = v.Phone
		template["Name"] = v.Name

		input, _ := json.Marshal(&template)

		resp, err := http.Post("https://waba-js-666dj.ondigitalocean.app/dev/broadcast", "Content-Type: application/json", bytes.NewReader(input))
		if err != nil {
			log.Println(err)
		}
		if resp.StatusCode == fiber.StatusOK {
			// post msg to message server
		} else {
			failedList = append(failedList, v)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"failedList": failedList})
}

func AddMessageTest(c *fiber.Ctx) error {
	msg := new(model.Message)

	err := c.BodyParser(&msg)
	if err != nil {
		log.Println("error in body parser, ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	URL := "http://message-server.message-server.svc.cluster.local:8080/api-v2/message"
	//URL := "http://localhost:8080/api-v2/message"
	uVal := url.Values{}
	uVal.Set("cid", "tiffany")
	uVal.Set("cname", "WABA")
	uVal.Set("rid", msg.RoomID)
	uVal.Set("ts", msg.Timestamp)
	URL = URL + "?" + uVal.Encode()

	item := c.Body()
	//item, _ := json.Marshal(&msg)

	resp, err := http.Post(URL, "Content-Type: application/json", bytes.NewBuffer(item))
	if err != nil {
		log.Println("error in post request", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	return c.Status(resp.StatusCode).JSON(data)
}
