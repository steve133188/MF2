package handler

import (
	"Move/config"
	"Move/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func MoveChatroom(c *fiber.Ctx) error {
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

	ctx := context.Background()
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Start to fill in Redis")

	for _, v := range chatrooms {
		temp := structs.Map(&v)
		temp["BotOn"] = strconv.FormatBool(v.BotOn)
		temp["IsPin"] = strconv.FormatBool(v.IsPin)
		err := RedisClient.HMSet(ctx, v.Channel+":Chatroom:"+v.RoomID, temp).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	return c.Status(fiber.StatusOK).JSON(chatrooms)

}

func MoveCustomers(c *fiber.Ctx) error {
	customers := make([]model.Customer, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("CUSTOMERTABLE")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pCustomers := make([]model.Customer, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pCustomers)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		customers = append(customers, pCustomers...)
	}

	ctx := context.Background()
	pipe := config.RedisClient.Pipeline()
	fmt.Println("Start to fill in Redis")

	for _, v := range customers {
		temp := structs.Map(&v)
		temp["Channels"], _ = json.Marshal(temp["Channels"])
		temp["AgentsID"], _ = json.Marshal(temp["AgentsID"])
		temp["TagsID"], _ = json.Marshal(temp["TagsID"])

		pipe.HMSet(ctx, "Customer:"+strconv.Itoa(v.CustomerID), temp)

	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(customers)

}

func MoveUsers(c *fiber.Ctx) error {
	users := make([]model.User, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("USERTABLE")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pUsers := make([]model.User, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pUsers)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		users = append(users, pUsers...)
	}

	ctx := context.Background()
	pipe := config.RedisClient.Pipeline()
	fmt.Println("Start to fill in Redis")

	for _, v := range users {
		temp := structs.Map(&v)
		temp["Channels"], _ = json.Marshal(temp["Channels"])
		temp["IsBot"] = strconv.FormatBool(v.IsBot)

		pipe.HMSet(ctx, "User:"+strconv.Itoa(v.UserID), temp)

	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(users)

}

func MoveTags(c *fiber.Ctx) error {
	tags := make([]model.Tag, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("TAGTABLE")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pTags := make([]model.Tag, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pTags)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		tags = append(tags, pTags...)
	}

	ctx := context.Background()
	pipe := config.RedisClient.Pipeline()
	fmt.Println("Start to fill in Redis")

	for _, v := range tags {
		temp := structs.Map(&v)

		pipe.HMSet(ctx, "Tag:"+strconv.Itoa(v.TagID), temp)

	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(tags)

}

func MoveTeams(c *fiber.Ctx) error {
	teams := make([]model.Team, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("ORGTABLE")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pTeams := make([]model.Team, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pTeams)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		teams = append(teams, pTeams...)
	}

	ctx := context.Background()
	pipe := config.RedisClient.Pipeline()
	fmt.Println("Start to fill in Redis")

	for _, v := range teams {
		temp := structs.Map(&v)
		temp["ChildrenID"], _ = json.Marshal(temp["ChildrenID"])

		pipe.HMSet(ctx, "Team:"+strconv.Itoa(v.TeamID), temp)

	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(teams)

}

func MoveRoles(c *fiber.Ctx) error {
	roles := make([]model.Role, 0)
	p := dynamodb.NewScanPaginator(config.DynaClient, &dynamodb.ScanInput{
		TableName: aws.String(config.GoDotEnvVariable("ROLETABLE")),
	})

	for p.HasMorePages() {
		outs, err := p.NextPage(c.Context())
		if err != nil {
			log.Println("Error in Scanning,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		pRoles := make([]model.Role, 0)
		err = attributevalue.UnmarshalListOfMaps(outs.Items, &pRoles)
		if err != nil {
			log.Println("Error in UnmarshalListOfMaps data,", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		roles = append(roles, pRoles...)
	}

	ctx := context.Background()
	pipe := config.RedisClient.Pipeline()
	fmt.Println("Start to fill in Redis")

	for _, v := range roles {
		temp := structs.Map(&v)
		temp["Auth"], _ = json.Marshal(temp["Auth"])

		pipe.HMSet(ctx, "Role:"+strconv.Itoa(v.RoleID), temp)

	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(roles)

}
