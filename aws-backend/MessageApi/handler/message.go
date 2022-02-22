package handler

import (
	"log"
	"mf2-message-api/config"
	"mf2-message-api/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

func GetAllMessagesByChatroom(c *fiber.Ctx) error {
	roomId := c.Params("roomId")

	messages := make([]model.Message, 0)

	p := dynamodb.NewQueryPaginator(config.DynaClient, &dynamodb.QueryInput{
		TableName:              aws.String(config.GoDotEnvVariable("MESSAGE")),
		KeyConditionExpression: aws.String("room_id = :rid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":rid": &types.AttributeValueMemberS{Value: roomId},
		},
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("[ERROR] message get all,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pMessages := make([]model.Message, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pMessages)
		if err != nil {
			log.Println("[ERROR] message get all,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		messages = append(messages, pMessages...)
	}

	if len(messages) == 0 {
		c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}

func AddMessage(c *fiber.Ctx) error {
	message := new(model.Message)

	err := c.BodyParser(&message)
	if err != nil {
		log.Println("[ERROR] add message unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(message.RoomID) == 0 || len(message.Timestamp) == 0 {
		log.Println("[ERROR] add message, missing timestamp or room_id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing timestamp or room_id"})
	}

	item, err := attributevalue.MarshalMap(&message)
	if err != nil {
		log.Println("[ERROR] add message marshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = config.DynaClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName:    aws.String(config.GoDotEnvVariable("MESSAGE")),
		Item:         item,
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		log.Println("[ERROR] add message,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
