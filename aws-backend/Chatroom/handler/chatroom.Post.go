package handler

import (
	"aws-lambda-chatroom/config"
	"aws-lambda-chatroom/model"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

func GetChatroomByAgent(c *fiber.Ctx) error {
	var data struct {
		UserID int `json:"user_id"`
		TeamID int `json:"team_id"`
		RoleID int `json:"role_id"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("error in unmarshal request body,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in unmarshal request body, " + err.Error()})
	}

	if data.RoleID == 0 || data.UserID == 0 || data.TeamID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing input data, " + err.Error()})
	}

	rout, err := config.DynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(config.GoDotEnvVariable("ROLETABLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.RoleID)},
		},
	})
	if err != nil {
		log.Println("error in get role item,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in get role item, " + err.Error()})
	}

	role := new(model.Role)
	err = attributevalue.UnmarshalMap(rout.Item, &role)
	if err != nil {
		log.Println("error in unmarshal role item,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in unmarshal role item, " + err.Error()})
	}

	var p *dynamodb.ScanPaginator
	if role.Auth.All {
		p = dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
			TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
		})
	} else {
		filterStr := ""
		filterExp := make(map[string]types.AttributeValue)
		if role.Auth.WABA {
			filterStr = "contains(channels, :waba) or "
			filterExp[":waba"] = &types.AttributeValueMemberS{Value: "WABA"}
		}
		if role.Auth.Whatsapp {
			filterStr += "team_id = :tid"
			filterExp[":tid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.TeamID)}
		} else {
			filterStr += "user_id = :uid"
			filterExp[":uid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)}
		}
		p = dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
			TableName:                 aws.String(config.GoDotEnvVariable("CHATROOM")),
			FilterExpression:          aws.String(filterStr),
			ExpressionAttributeValues: filterExp,
		})
		log.Println("filterStr = ", filterStr)
	}

	chatrooms := make([]model.Chatroom, 0)

	count := 0
	for p.HasMorePages() {
		souts, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("error in scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in scanning, " + err.Error()})
		}

		pChatrooms := make([]model.Chatroom, 0)
		err = attributevalue.UnmarshalListOfMaps(souts.Items, &pChatrooms)
		if err != nil {
			fmt.Println("error in unmarshal data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error in unmarshal data, " + err.Error()})
		}
		chatrooms = append(chatrooms, pChatrooms...)
		count += int(souts.Count)
	}
	log.Println("count = ", count)

	return c.Status(fiber.StatusOK).JSON(chatrooms)
}
