package handler

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamodbConfig() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("REGION")
		return nil
	})
	if err != nil {
		log.Println("AddItem LoadDefaultConfig    ", err)
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc, nil

}
