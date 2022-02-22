package handler

import (
	"log"
	"mf2-message-api/config"
	"mf2-message-api/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
)

func AddActivity(c *fiber.Ctx) error {
	acty := new(model.Activity)

	err := c.BodyParser(&acty)
	if err != nil {
		log.Println("[ERROR] add activity")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	item, err := attributevalue.MarshalMap(&acty)
	if err != nil {
		log.Println("[ERROR] add activity")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = config.DynaClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName: aws.String(config.GoDotEnvVariable("ACTIVITY")),
		Item:      item,
	})
	if err != nil {
		log.Println("[ERROR] add activity")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
