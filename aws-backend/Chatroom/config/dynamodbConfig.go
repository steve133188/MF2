package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynaClient *dynamodb.Client

func DynamodbConfig() {
	AWS_ACCESS_KEY_ID := GoDotEnvVariable("KEYID1") + GoDotEnvVariable("KEYID2")
	AWS_SECRET_ACCESS_KEY := GoDotEnvVariable("ACCESSKEY1") + GoDotEnvVariable("ACCESSKEY2")

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = GoDotEnvVariable("REGION")
		o.Credentials = credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     AWS_ACCESS_KEY_ID,
				SecretAccessKey: AWS_SECRET_ACCESS_KEY,
			},
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	DynaClient = dynamodb.NewFromConfig(cfg)
}
