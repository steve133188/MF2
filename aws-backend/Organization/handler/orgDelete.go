package handler

import (
	"aws-lambda-org/model"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteOrgItem(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	orgId := req.PathParameters["id"]
	if len(orgId) == 0 {
		log.Println("MissingOrgID")
		return ApiResponse(http.StatusBadRequest, ErrMsg{aws.String("MissingOrgID")}), nil
	}

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: orgId},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, OrgID = ", orgId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem, OrgID = " + orgId + ", " + err.Error())}), nil
	}
	if out.Item == nil {
		fmt.Println("OrgUnitNotexists, OrgID = ", orgId, ", ", err)
		return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("OrgUnitNotexists, OrgID = " + orgId + ", " + err.Error())}), nil
	}

	org := new(model.Organization)
	err = attributevalue.UnmarshalMap(out.Item, &org)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, OrgID = ", orgId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, OrgID = " + orgId + ", " + err.Error())}), nil
	}

	orgIdInt, err := strconv.Atoi(orgId)
	if err != nil {
		fmt.Println("FailedToConvert, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToConvert, " + err.Error())}), nil
	}

	if org.ParentID != 0 {
		out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
			TableName: aws.String(table),
			Key: map[string]types.AttributeValue{
				"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(org.ParentID)},
			},
		})
		if err != nil {
			fmt.Println("FailedToGetItem, ParentID = ", org.ParentID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToGetItem, ParentID = " + strconv.Itoa(org.ParentID) + orgId + ", " + err.Error())}), nil
		}
		if out.Item == nil {
			fmt.Println("ParentItemNotExist, ParentID = ", org.ParentID, ", ", err)
			return ApiResponse(http.StatusNotFound, ErrMsg{aws.String("ParentItemNotExist, ParentID = " + strconv.Itoa(org.ParentID) + orgId + ", " + err.Error())}), nil
		}
		parent := new(model.Organization)
		err = attributevalue.UnmarshalMap(out.Item, &parent)
		if err != nil {
			fmt.Println("FailedToUnmarshalMap, ParentID = ", org.ParentID, ", ", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUnmarshalMap, ParentID = " + strconv.Itoa(org.ParentID) + orgId + ", " + err.Error())}), nil
		}

		for k, v := range parent.ChildrenID {
			if v == orgIdInt {
				delStr := "REMOVE children_id[" + strconv.Itoa(k) + "]"
				_, err := dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
					TableName: aws.String(table),
					Key: map[string]types.AttributeValue{
						"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(org.ParentID)},
					},
					UpdateExpression: aws.String(delStr),
				})
				if err != nil {
					fmt.Println("FailedToUpdateParent, ParentID = ", org.ParentID, ", ", err)
					return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToUpdateParent, ParentID = " + strconv.Itoa(org.ParentID) + orgId + ", " + err.Error())}), nil
				}
				break
			}
		}

	}

	idStr, err := strconv.Atoi(orgId)
	if err != nil {
		fmt.Println("FailedToConvertOrgID, OrgID = ", orgId, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToConvertOrgID, OrgID = " + orgId + ", " + err.Error())}), nil
	}

	id, err := DeleteChildrenUnits(dynaClient, idStr, table)
	if err != nil {
		fmt.Println("FailedToDeleteItem, OrgID = ", id, ", ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToDeleteItem, OrgID = " + orgId + ", " + err.Error())}), nil
	}

	return ApiResponse(http.StatusOK, map[string]string{"message": "success"}), nil
}

func DeleteChildrenUnits(dynaClient *dynamodb.Client, id int, table string) (int, error) {
	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, OrgID = ", id, ", ", err)
		return id, err
	}

	if out.Item == nil {
		fmt.Println("OrgUnitNotExisted, OrgID = ", id)
		return id, errors.New("OrgUnitNotExisted")
	}

	org := new(model.Organization)

	err = attributevalue.UnmarshalMap(out.Item, &org)
	if err != nil {
		fmt.Println("FailedToUnmarshalMap, OrgID = ", id, ", ", err)
		return id, err
	}

	if len(org.ChildrenID) != 0 {
		for _, v := range org.ChildrenID {
			i, err := DeleteChildrenUnits(dynaClient, v, table)
			if err != nil {
				return i, err
			}
		}
	}

	_, err = dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	})
	if err != nil {
		fmt.Println("FailedToDeleteItem, OrgID = ", id, ", ", err)
		return id, err
	}

	puser := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("USERTABLE")),
		FilterExpression: aws.String("team_id = :tid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":tid": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
		Limit: aws.Int32(50),
	})

	users := make([]model.User, 0)
	for puser.HasMorePages() {
		uouts, err := puser.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScan, OrgID = ", id)
			return 0, err
		}
		pusers := make([]model.User, 0)

		err = attributevalue.UnmarshalListOfMaps(uouts.Items, &pusers)
		if err != nil {
			fmt.Println("FailedToUnmarshalListOfMaps, OrgID = ", id)
			return 0, err
		}

		users = append(users, pusers...)
	}

	for _, v := range users {
		_, err = dynaClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(os.Getenv("USERTABLE")),
			Key: map[string]types.AttributeValue{
				"user_id": &types.AttributeValueMemberN{Value: strconv.Itoa(v.UserID)},
			},
			UpdateExpression: aws.String("SET team_id = :tid"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":tid": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
			},
		})
		if err != nil {
			fmt.Println("FailedToUpdateItem, UserID = ", v.UserID)
			return 0, err
		}
	}

	return id, nil
}
