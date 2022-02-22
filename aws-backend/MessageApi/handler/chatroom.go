package handler

import (
	"context"
	"fmt"
	"log"
	"mf2-message-api/config"
	"mf2-message-api/model"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

func GetChatroomsByAgent(c *fiber.Ctx) error {
	var data struct {
		UserID int `json:"user_id"`
		TeamID int `json:"team_id"`
		RoleID int `json:"role_id"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("[ERROR] chatroom get by agent,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	rout, err := config.DynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(config.GoDotEnvVariable("ROLE")),
		Key: map[string]types.AttributeValue{
			"role_id": &types.AttributeValueMemberN{Value: strconv.Itoa(data.RoleID)},
		},
	})
	if err != nil {
		log.Println("[ERROR] chatroom get role,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	role := new(model.Role)
	err = attributevalue.UnmarshalMap(rout.Item, &role)
	if err != nil {
		log.Println("[ERROR] chatroom unmarshal role,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var p *dynamodb.ScanPaginator
	if role.Auth.All {
		p = dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
			TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
		})
	} else {
		filterStr := "user_id = :uid"
		filterExp := make(map[string]types.AttributeValue)
		filterExp[":uid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.UserID)}
		if role.Auth.WABA {
			filterStr = " or contains(channels, :waba)"
			filterExp[":waba"] = &types.AttributeValueMemberS{Value: "WABA"}
		}
		if role.Auth.Whatsapp {
			filterStr += " or team_id = :tid"
			filterExp[":tid"] = &types.AttributeValueMemberN{Value: strconv.Itoa(data.TeamID)}
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
			fmt.Println("[ERROR] chatroom scan,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pChatrooms := make([]model.Chatroom, 0)
		err = attributevalue.UnmarshalListOfMaps(souts.Items, &pChatrooms)
		if err != nil {
			fmt.Println("[ERROR] chatroom unmarshal,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		chatrooms = append(chatrooms, pChatrooms...)
		count += int(souts.Count)
	}
	log.Println("count = ", count)

	return c.Status(fiber.StatusOK).JSON(chatrooms)
}

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
		log.Println("[ERROR] chatroom update,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	chatroom := new(model.Chatroom)
	err = attributevalue.UnmarshalMap(out.Attributes, &chatroom)
	if err != nil {
		log.Println("[ERROR] chatroom update unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}

func CheckChatroom(c *fiber.Ctx) error {
	chatroom := new(model.Chatroom)

	channel := c.Params("channel")
	roomId := c.Params("room_id")

	out, err := config.DynaClient.GetItem(c.Context(), &dynamodb.GetItemInput{
		TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
		Key: map[string]types.AttributeValue{
			"channel": &types.AttributeValueMemberS{Value: channel},
			"room_id": &types.AttributeValueMemberS{Value: roomId},
		},
	})
	if err != nil {
		log.Println("[ERROR] chatroom check get item, ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(out.Item) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = attributevalue.UnmarshalMap(out.Item, &chatroom)
	if err != nil {
		log.Println("[ERROR] chatroom check unmarshal, ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}

func AddChatroom(c *fiber.Ctx) error {
	chatroom := new(model.Chatroom)

	err := c.BodyParser(&chatroom)
	if err != nil {
		log.Println("[ERROR] add chatroom unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(chatroom.Channel) == 0 || len(chatroom.RoomID) == 0 {
		log.Println("[ERROR] add chatroom, missing channel or room_id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing channel or room_id"})
	}

	item, err := attributevalue.MarshalMap(&chatroom)
	if err != nil {
		log.Println("[ERROR] add chatroom marshal, ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	out, err := config.DynaClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName:    aws.String(config.GoDotEnvVariable("CHATROOM")),
		Item:         item,
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		log.Println("[ERROR] add chatroom,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &chatroom)
	if err != nil {
		log.Println("[ERROR] add chatroom unmarshalmap,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}

func UpdateChatroomName(c *fiber.Ctx) error {
	chatroom := new(model.Chatroom)

	err := c.BodyParser(chatroom)
	if err != nil {
		log.Println("[ERROR] update chatroom unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(chatroom.Channel) == 0 || len(chatroom.RoomID) == 0 {
		log.Println("[ERROR] update chatroom, missing channel or room_id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing channel or room_id"})
	}

	out, err := config.DynaClient.UpdateItem(c.Context(), &dynamodb.UpdateItemInput{
		TableName: aws.String(config.GoDotEnvVariable("CHATROOM")),
		Key: map[string]types.AttributeValue{
			"channel": &types.AttributeValueMemberS{Value: chatroom.Channel},
			"room_id": &types.AttributeValueMemberS{Value: chatroom.RoomID},
		},
		UpdateExpression: aws.String("set #n = :n"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{Value: chatroom.Name},
		},
		ExpressionAttributeNames: map[string]string{
			"#n": "name",
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		log.Println("[ERROR] update chatroom,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &chatroom)
	if err != nil {
		log.Println("[ERROR] update unmarshalmap,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(chatroom)
}
