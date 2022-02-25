package handler

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"mf2-channel-router/config"
	"mf2-channel-router/model"
	"net/http"
	"net/url"
	"strings"
)

func ChannelConnect(c *fiber.Ctx) error {
	connectData := new(model.Connect)

	err := c.BodyParser(&connectData)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	hKey := connectData.CID + ":channels:" + connectData.CName + ":" + connectData.UID

	err = config.ClusterClient.HSet(c.Context(), hKey, "status", "CONNECTING").Err()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	redisData, err := config.ClusterClient.HGetAll(c.Context(), hKey).Result()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// node index
	NI := redisData["node_name"]
	URL := redisData["url"]

	resource := "/connect"
	data := url.Values{}
	data.Set("user_id", connectData.UID)
	data.Set("node_index", NI)
	data.Set("user_name", connectData.UName)
	data.Set("team_id", connectData.TID)

	urlStr := URL + resource
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	return c.Status(fiber.StatusOK).JSON(resp.Body)
}

func ChannelRestart(c *fiber.Ctx) error {
	cid := c.Query("cid")
	cname := c.Query("cname")
	uid := c.Query("uid")

	hKey := cid + ":channels:" + cname + ":" + uid

	url, err := config.ClusterClient.HGet(c.Context(), hKey, "url").Result()
	if err != nil {
		log.Println("HGetAll error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	url += "/restart"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("restart error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()
	return c.SendStatus(fiber.StatusOK)
}

func ChannelDisconnect(c *fiber.Ctx) error {
	cid := c.Query("cid")
	cname := c.Query("cname")
	uid := c.Query("uid")

	hKey := cid + ":channels:" + cname + ":" + uid

	url, err := config.ClusterClient.HGet(c.Context(), hKey, "url").Result()
	if err != nil {
		log.Println("HGetAll error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	url += "/disconnect"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("restart error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()
	return c.SendStatus(fiber.StatusOK)
}

func ChannelSendMessage(c *fiber.Ctx) error {
	cid := c.Query("cid")
	cname := c.Query("cname")
	uid := c.Query("uid")

	hKey := cid + ":channels:" + cname + ":" + uid
	URL, err := config.ClusterClient.HGet(c.Context(), hKey, "url").Result()
	if err != nil {
		log.Println("HGetAll error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	URL += "/send-message"

	//message := new(model.Message)
	//err = c.BodyParser(&message)
	//if err != nil {
	//	log.Println("parser error,", err)
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	//data := url.Values{}
	//data.Set("room_id", message.RoomID)
	//data.Set("timestamp", message.Timestamp)
	//data.Set("status", message.Status)
	//data.Set("message_type", message.MessageType)
	//data.Set("hasQuotedMsg", strconv.FormatBool(message.HasQuotedMsg))
	//data.Set("is_media", strconv.FormatBool(message.IsMedia))
	//data.Set("message_id", message.MessageID)
	//data.Set("channel", message.Channel)
	//data.Set("media_id", message.MediaUrl)
	//data.Set("sender", message.Sender)
	//data.Set("recipient", message.Recipient)
	//data.Set("read", strconv.FormatBool(message.Read))
	//data.Set("is_forwarded", strconv.FormatBool(message.IsForwarded))
	//data.Set("from_me", strconv.FormatBool(message.FromMe))
	//data.Set("link", message.Link)
	//data.Set("body", message.Body)
	//data.Set("quote", message.Quote)

	inputData := c.Body()

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(inputData))
	if err != nil {
		log.Println("request error,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("do request error,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body error,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(resp.StatusCode).JSON(body)
}

func ChannelUpdateStatus(c *fiber.Ctx) error {
	cid := c.Query("cid")
	cname := c.Query("cname")
	uid := c.Query("uid")
	status := c.Query("status")

	hKey := cid + ":channels:" + cname + ":" + uid

	err := config.ClusterClient.HSet(c.Context(), hKey, "status", status).Err()
	if err != nil {
		log.Println("hset error,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
