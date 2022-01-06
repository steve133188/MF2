package handler

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func getUserLeads(userId string, dynaClient *dynamodb.Client) (int, error) {
	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("CUSTOMERTABLE")),
		Limit:            aws.Int32(80),
		FilterExpression: aws.String("contains(agents_id, :id)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: userId},
		},
	})
	count := 0
	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToGetLeads, UserID = ", userId)
			return 0, errors.New("FailedToGetLeads, UserID = " + userId)
		}

		count += int(out.Count)
	}
	return count, nil
}
