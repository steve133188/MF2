package Services

import (
	"aws-lambda-chatroom/Handler"
	"aws-lambda-chatroom/Model"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ErrMsg struct {
	Error *string `json:"error"`
}

func GetItems(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	chatroomID := req.QueryStringParameters["room_id"]
	userID := req.QueryStringParameters["user_id"]
	if req.Path == "/prod/chatroom" {
		if len(chatroomID) != 0 || len(userID) != 0 {
			chatroom := new(Model.ChatRoom)
			data, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
				TableName: aws.String(tableName),
				Key: map[string]types.AttributeValue{
					"room_id": &types.AttributeValueMemberN{Value: chatroomID},
					"user_id": &types.AttributeValueMemberN{Value: userID},
				},
			})
			if err != nil {
				log.Println("GetItem ErrorToGetItem chatroomID: ", chatroomID, " ", err)
				return Handler.ApiResponse(http.StatusOK, ErrMsg{aws.String("FailedTOGetItem")}), nil

			}

			if data.Item == nil {
				log.Println("GetItem ItemNotExisted:    ", err)
				return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ItemNotExisted")}), nil

			}

			err = attributevalue.UnmarshalMap(data.Item, &chatroom)
			if err != nil {
				log.Println("GetItem ErrorToUnmarshalMap:    ", err)
				return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ErrorToUnmarshalMap")}), nil

			}

			return Handler.ApiResponse(http.StatusOK, chatroom), nil
		} else {
			return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingChatroomIdORUserId")}), nil
		}
	} else if req.Path == "/prod/chatrooms" {
		p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
			TableName: aws.String(tableName),
			Limit:     aws.Int32(1),
		})

		var chatrooms []Model.ChatRoom = make([]Model.ChatRoom, 0)

		for p.HasMorePages() {
			out, err := p.NextPage(context.TODO())
			if err != nil {
				log.Println("GetItems  FailedToScanNextPage    ", err)
				return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToScanNextPage")}), err
			}

			var pChatrooms []Model.ChatRoom = make([]Model.ChatRoom, 0)
			err = attributevalue.UnmarshalListOfMaps(out.Items, &pChatrooms)
			if err != nil {
				log.Println("GetItems  FailedToUnmarshalListOfMaps    ", err)
				return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalListOfMaps")}), err
			}

			chatrooms = append(chatrooms, pChatrooms...)

		}

		return Handler.ApiResponse(http.StatusInternalServerError, chatrooms), nil
	} else {
		return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("InvalidMethodOrRoute")}), nil

	}

}

func AddItem(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {

	chatroom := new(Model.ChatRoom)
	err := json.Unmarshal([]byte(req.Body), &chatroom)
	if err != nil {
		log.Println("AddItem Unmarshal    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalData")}), nil
	}

	if chatroom.UserID == 0 {
		log.Println("AddItem Missing UserID")
		return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingUserId")}), nil
	}

	if chatroom.RoomID == 0 {
		rand.Seed(time.Now().UnixNano())
		id := rand.Intn(99999999)
		chatroom.RoomID = id
	}

	av, err := attributevalue.MarshalMap(&chatroom)
	if err != nil {
		log.Println("AddItem MarshalMap    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMapData")}), nil

	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(room_id) AND attribute_not_exists(user_id)"),
	})

	if err != nil {
		log.Println("AddItem PutItem    ", err)
		return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("ExistedChatroomIDAndUserID")}), nil

	}

	return Handler.ApiResponse(http.StatusCreated, chatroom), nil
}

func EditItem(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	chatroom := new(Model.ChatRoom)

	err := json.Unmarshal([]byte(req.Body), &chatroom)
	if err != nil {
		log.Println("EditItem Unmarshal    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalData")}), err
	}

	av, err := attributevalue.MarshalMap(&chatroom)
	if err != nil {
		log.Println("EditItem MarshalMap    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToMarshalMapData")}), err

	}

	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		log.Println("EditItem PutItem    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateItem")}), nil

	}

	return Handler.ApiResponse(http.StatusOK, chatroom), nil

}

func DeleteItem(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	chatroomID := req.QueryStringParameters["room_id"]
	userID := req.QueryStringParameters["user_id"]

	if len(chatroomID) == 0 || len(userID) == 0 {
		log.Println("DeleteItem MissingChatroomIdOrUserId")
		return Handler.ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingChatroomIdOrUserId")}), nil
	}

	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"room_id": &types.AttributeValueMemberN{Value: chatroomID},
			"user_id": &types.AttributeValueMemberN{Value: userID},
		},
		ConditionExpression: aws.String("attribute_exists(room_id) AND attribute_exists(user_id)"),
	})
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			log.Println("DeleteItem ItemNotExisted    ", err)
			return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ItemNotExisted")}), nil
		}
		log.Println("DeleteItem DeleteItem    ", err)
		return Handler.ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem")}), nil
	}
	return Handler.ApiResponse(http.StatusOK, map[string]string{"message": "success"}), nil

}
