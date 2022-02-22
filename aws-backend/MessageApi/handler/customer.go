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

func AddCustomer(c *fiber.Ctx) error {
	customer := new(model.Customer)

	err := c.BodyParser(&customer)
	if err != nil {
		log.Println("[ERROR] customer add unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	item, err := attributevalue.MarshalMap(&customer)
	if err != nil {
		log.Println("[ERROR] customer add marshalmap,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	out, err := config.DynaClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName:    aws.String(config.GoDotEnvVariable("CUSTOMER")),
		Item:         item,
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		log.Println("[ERROR] customer add,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = attributevalue.UnmarshalMap(out.Attributes, &customer)
	if err != nil {
		log.Println("[ERROR] customer add unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func GetCustomer(c *fiber.Ctx) error {
	customer := new(model.Customer)

	id := c.Params("id")

	out, err := config.DynaClient.GetItem(c.Context(), &dynamodb.GetItemInput{
		TableName: aws.String(config.GoDotEnvVariable("CUSTOMER")),
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		log.Println("[ERROR] customer get,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(out.Item) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = attributevalue.UnmarshalMap(out.Item, &customer)
	if err != nil {
		log.Println("[ERROR] customer get unmarshal,", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
