package handler

import (
	"aws-message-api/config"
	"aws-message-api/model"
	"log"

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
		TableName:              aws.String(config.GoDotEnvVariable("MESSAGETABLE")),
		KeyConditionExpression: aws.String("room_id = :rid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":rid": &types.AttributeValueMemberS{Value: roomId},
		},
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Query,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in Query, " + err.Error()})
		}

		pMessages := make([]model.Message, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pMessages)
		if err != nil {
			log.Println("Error in Unmarshal list of data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in Unmarshal list of data, " + err.Error()})
		}

		messages = append(messages, pMessages...)
	}

	if len(messages) == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No Data found with room ID = " + roomId})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
