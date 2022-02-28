package hanlder

import (
	"ChatbotAPI/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateFlow(c *fiber.Ctx) error {

	var data map[string]interface{}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("error in unmarshal request body,", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error in unmarshal request body, " + err.Error()})
	}

	flowKey := fmt.Sprintf("%v:flow:%v", data["companyId"], data["flowName"])
	fmt.Println(flowKey)

	isExist, err := config.ChatBotDB.Exists(context.Background(), flowKey).Result()
	if err != nil {
		log.Println("error in exist checking,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in exist checking, " + err.Error()})
	}
	if isExist > 0 {
		log.Println("data already exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "data already exist"})
	}

	marshalData, err := json.Marshal(data)
	if err != nil {
		log.Println("error in marshal request body,", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error in marshal request body, " + err.Error()})
	}

	err = config.ChatBotDB.Set(context.Background(), flowKey, marshalData, 0).Err()
	if err != nil {
		log.Println("error in creating data,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in creating data, " + err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func CreateOption(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}

func CreateAction(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}

//func GenerateFlowKey(companyId string) string {
//	return companyId + ":flows:" + strconv.FormatInt(time.Now().UnixMicro(), 10)
//}

func GenerateOptionKey(companyId string, stage string, parentKey string, condition []string) string {
	var temp string
	for _, v := range condition {
		temp += "#" + v
	}
	return companyId + ":automations:" + stage + ":" + parentKey + temp
}

//func GenerateActionKey(companyId string) string {
//	return companyId + ":actions:" + strconv.FormatInt(time.Now().UnixMicro(), 10)
//}

//func CreateChatList(c *fiber.Ctx) error {
//	var data map[string]interface{}
//
//	err := c.BodyParser(&data)
//	if err != nil {
//		log.Println("error in unmarshal request body,", err)
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error in unmarshal request body, " + err.Error()})
//	}
//
//	isExist, err := config.ChatBotDB.Exists(context.Background(), data["FlowID"]).Result()
//	if err != nil {
//		log.Println("error in exist checking,", err)
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in exist checking, " + err.Error()})
//	}
//	if isExist > 0 {
//		log.Println("data already exist")
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "data already exist"})
//	}
//
//	err = config.ChatBotDB.HMSet(context.Background(), data["FlowID"], data).Err()
//	if err != nil {
//		log.Println("error in creating data,", err)
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in creating data, " + err.Error()})
//	}
//
//	return c.SendStatus(fiber.StatusCreated)
//}
