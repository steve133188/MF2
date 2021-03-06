package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"mf2-message-server/config"
)

func GetAllMessagesByChatroom(c *fiber.Ctx) error {
	return nil
}

func AddMessage(c *fiber.Ctx) error {
	cid := c.Query("cid")
	cname := c.Query("cname")
	rid := c.Query("rid")
	ts := c.Query("ts")

	hKey := cid + ":messages:" + cname + ":" + rid + ":" + ts

	msg := make(map[string]interface{})
	body := c.Body()
	err := json.Unmarshal(body, &msg)
	if err != nil {
		log.Println("error in unmarshal", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	//err = config.TestClient.HMSet(c.Context(), hKey, msg).Err()
	err = config.ClusterClient.HMSet(c.Context(), hKey, msg).Err()
	if err != nil {
		log.Println("error in HMSet", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	pubMsg := make(map[string]interface{})
	pubMsg[hKey] = msg
	item, _ := json.Marshal(pubMsg)
	err = config.ClusterClient.Publish(c.Context(), "messages", item).Err()
	if err != nil {
		log.Println("error in publish", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("published message")
	if msg["from_me"] == false {
		pubMsg := make(map[string]interface{})
		pubMsg[hKey] = msg
		item, _ := json.Marshal(pubMsg)
		err = config.ClusterClient.Publish(c.Context(), "messages.received.WABA", item).Err()
		if err != nil {
			log.Println("error in publish", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		log.Println("published message")
	}

	return c.SendStatus(fiber.StatusOK)
}
