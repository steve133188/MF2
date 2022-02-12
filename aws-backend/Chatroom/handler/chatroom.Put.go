package handler

import (
	"aws-lambda-chatroom/config"
	"aws-lambda-chatroom/model"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

func UpdateChatroomunreadToZero(c *fiber.Ctx) error {
	pk := c.Params("channel")
	sk := c.Params("room_id")

	out, err := config.DynaClient.UpdateItem(c.Context(), &dynamodb.UpdateItemInput{
		TableName:           aws.String(config.GoDotEnvVariable("CHATROOM")),
		ConditionExpression: aws.String("attribute_exists(channel) AND attribute_exists(room_id)"),
		Key: map[string]types.AttributeValue{
			"channel": &types.AttributeValueMemberS{Value: pk},
			"room_id": &types.AttributeValueMemberS{Value: sk},
		},
		UpdateExpression: aws.String("SET #unread = :zero"),
		ExpressionAttributeNames: map[string]string{
			"#unread": "unread",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":zero": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		log.Println("Error in update item,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in update item, " + err.Error()})
	}

	chatroom := new(model.Chatroom)
	err = attributevalue.UnmarshalMap(out.Attributes, &chatroom)
	if err != nil {
		log.Println("Error in Unmarshal data,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in Unmarshal data, " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}
