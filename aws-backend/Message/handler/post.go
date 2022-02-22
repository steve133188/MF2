package handler

import (
	"aws-message-api/config"
	"aws-message-api/model"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
)

func AddMessage(c *fiber.Ctx) error {
	message := new(model.Message)

	err := c.BodyParser(&message)
	if err != nil {
		log.Println("error in unmarshal request body,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in unmarshal request body, " + err.Error()})
	}

	item, err := attributevalue.MarshalMap(&message)
	if err != nil {
		log.Println("error in marshal request body,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in marshal request body, " + err.Error()})
	}

	_, err = config.DynaClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName: aws.String(config.GoDotEnvVariable("MESSAGETABLE")),
		Item:      item,
	})
	if err != nil {
		log.Println("error in PutItem,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in PutItem, " + err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
