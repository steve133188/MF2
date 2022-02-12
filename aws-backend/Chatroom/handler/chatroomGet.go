package handler

import (
	"aws-lambda-chatroom/config"
	"aws-lambda-chatroom/model"
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

func GetChatrooms(c *fiber.Ctx) error {

	chatrooms := make([]model.Chatroom, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pChatrooms := make([]model.Chatroom, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pChatrooms)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		chatrooms = append(chatrooms, pChatrooms...)
	}

	return c.Status(fiber.StatusOK).JSON(chatrooms)
}

func GetOneChatroom(c *fiber.Ctx) error {
	pk := c.Params("channel")
	sk := c.Params("room_id")

	if len(pk) == 0 || len(sk) == 0 {
		log.Println("missing partition key or sort key")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing partition key or sort key"})
	}

	out, err := config.DynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
		Key: map[string]types.AttributeValue{
			"channel": &types.AttributeValueMemberS{Value: pk},
			"room_id": &types.AttributeValueMemberS{Value: sk},
		},
	})
	if err != nil {
		log.Println("Error in Get Item", pk, sk, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("Error in Query" + pk + sk + err.Error())})
	}

	chatroom := new(model.Chatroom)
	err = attributevalue.UnmarshalMap(out.Item, &chatroom)
	if err != nil {
		log.Println("Error in UnmarshalMap", pk, sk, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("Error in UnmarshalMap" + pk + sk + err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}

func GetChatroomsByUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	if len(userId) == 0 {
		log.Println("missing partition user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing partition user ID"})
	}

	chatrooms := make([]model.Chatroom, 0)

	p := dynamodb.NewQueryPaginator(config.DynaClient, &dynamodb.QueryInput{
		TableName:              aws.String(config.GoDotEnvVariable("CHATROOM")),
		IndexName:              aws.String("user_id-index"),
		KeyConditionExpression: aws.String("user_id = :user_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_id": &types.AttributeValueMemberN{Value: userId},
		},
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Query,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pChatrooms := make([]model.Chatroom, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pChatrooms)
		if err != nil {
			log.Println("Error in Unmarshal data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		chatrooms = append(chatrooms, pChatrooms...)
	}

	if len(chatrooms) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "no chatroom found with user ID = " + userId})
	}

	return c.Status(fiber.StatusOK).JSON(chatrooms)
}
