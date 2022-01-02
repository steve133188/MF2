package handler

import (
	"aws-lambda-org/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetOrgItemByID(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	orgId := req.PathParameters["id"]

	orgStructs := make([]model.OrgStruct, 0)

	id, err := strconv.Atoi(orgId)
	if err != nil {
		fmt.Println("FailedToConvert, ", err)
		return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("FailedToConvert" + err.Error())}), nil
	}

	results, err, status := GetChildrenOrgObject(dynaClient, id, table)
	if err != nil {
		fmt.Println("FailedToGetChildren, ", err)
		return ApiResponse(status, ErrMsg{aws.String(err.Error())}), nil
	}

	orgStructs = append(orgStructs, results...)

	return ApiResponse(http.StatusOK, orgStructs), nil
}

func GetOrgItems(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	orgs := make([]model.Organization, 0)

	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(100),
		FilterExpression: aws.String("parent_id = :parentID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":parentID": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
		},
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScanNextPage", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ScanningError")}), nil
		}

		var pageOrgs []model.Organization = make([]model.Organization, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pageOrgs)
		if err != nil {
			fmt.Println("UnmarshalListOfMapsError", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalListOfMapsError")}), nil
		}

		orgs = append(orgs, pageOrgs...)

	}

	orgStructs := make([]model.OrgStruct, 0)

	for _, v := range orgs {
		id := v.OrgID

		results, err, status := GetChildrenOrgObject(dynaClient, id, table)
		if err != nil {
			fmt.Println("FailedToGetChildren, ", err)
			return ApiResponse(status, ErrMsg{aws.String(err.Error())}), nil
		}

		orgStructs = append(orgStructs, results...)
	}

	return ApiResponse(http.StatusOK, orgStructs), nil

}

func GetChildrenOrgObject(dynaClient *dynamodb.Client, id int, table string) ([]model.OrgStruct, error, int) {
	org := new(model.Organization)

	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"org_id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	})
	if err != nil {
		fmt.Println("FailedToGetItem, OrgId = ", id)
		return nil, errors.New("FailedToGetItem, OrgId = " + strconv.Itoa(id)), http.StatusInternalServerError
	}
	if out.Item == nil {
		fmt.Println("OrgUnitNotExisted, OrgID = ", id)
		return nil, errors.New("OrgUnitNotExisted, OrgId = " + strconv.Itoa(id)), http.StatusInternalServerError
	}

	err = attributevalue.UnmarshalMap(out.Item, &org)
	if err != nil {
		fmt.Println("FailedToUNmarshalMap, OrgID = ", id)
		return nil, errors.New("OrgUnitNotExisted, OrgId = " + strconv.Itoa(id)), http.StatusInternalServerError
	}

	data := new(model.OrgStruct)
	err = attributevalue.UnmarshalMap(out.Item, &data)
	if err != nil {
		fmt.Println("FailedToUNmarshalMap, OrgID = ", id)
		return nil, errors.New("OrgUnitNotExisted, OrgId = " + strconv.Itoa(id)), http.StatusInternalServerError
	}

	results := make([]model.OrgStruct, 0)

	if len(org.ChildrenID) != 0 {
		data.Children = make([]model.OrgStruct, 0)
		for _, v := range org.ChildrenID {
			result, err, status := GetChildrenOrgObject(dynaClient, v, table)
			if err != nil {
				return nil, err, status
			}
			data.Children = append(data.Children, result...)
		}
	}

	results = append(results, *data)
	return results, nil, http.StatusOK
}

func GetTeamName(req events.APIGatewayProxyRequest, table string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	org := make([]model.Organization, 0)
	p := dynamodb.NewScanPaginator(dynaClient, &dynamodb.ScanInput{
		TableName:        aws.String(table),
		Limit:            aws.Int32(100),
		FilterExpression: aws.String("#type = :team"),
		ExpressionAttributeNames: map[string]string{
			"#type": "type",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":team": &types.AttributeValueMemberS{Value: "team"},
		},
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			fmt.Println("FailedToScanNextPage", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("ScanningError")}), nil
		}

		pOrg := make([]model.Organization, 0)
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pOrg)
		if err != nil {
			fmt.Println("UnmarshalListOfMapsError", err)
			return ApiResponse(http.StatusInternalServerError, ErrMsg{aws.String("UnmarshalListOfMapsError")}), nil
		}
		org = append(org, pOrg...)
	}

	return ApiResponse(http.StatusOK, org), nil
}
