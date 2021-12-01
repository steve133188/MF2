package Services

import (
	"context"
	"log"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//create one ORG
//find one org by id
//find

func CreateDivision(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	data := new(Model.ORG)

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("CreateDivision parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"parent_id", data.ParentID}, {"name", data.Name}}
	found, err := col.CountDocuments(c.Context(), filter, options.Count().SetLimit(1))
	if found > 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	data.ID = xid.New().String()
	if data.Type != "team" || data.Type != "division" {
		c.SendStatus(fiber.StatusBadRequest)
	}

	if data.ChildrenID == nil {
		data.ChildrenID = make([]string, 0)
	}

	if data.ParentID != "" {
		nFilter := bson.D{{"id", data.ParentID}}
		nUpdate := bson.D{{"$push", bson.D{{"children_id", data.ID}}}}
		_, err := col.UpdateOne(c.Context(), nFilter, nUpdate)
		if err != nil {
			log.Println("CreateDivision UpdateOne ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

	}

	result, err := col.InsertOne(c.Context(), data)
	if err != nil {
		log.Println("CreateDivision InsertOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = col.FindOne(c.Context(), bson.D{{"_id", result.InsertedID}}).Decode(&data)
	if err != nil {
		log.Println("CreateDivision FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetRootDivisions(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	var datas []Model.ORG = make([]Model.ORG, 0)

	filter := bson.D{{"parent_id", ""}}

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetRootDivision Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	err = cursor.All(c.Context(), &datas)
	if err != nil {
		log.Println("GetRootDivision All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(datas)
}

func GetOrgByParentID(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	parentID := c.Params("parentId")

	filter := bson.D{{"parent_id", parentID}}

	var datas []Model.ORG = make([]Model.ORG, 0)
	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetOrgByParentID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	err = cursor.All(c.Context(), &datas)
	if err != nil {
		log.Println("GetOrgByParentID All ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(datas)
}

func GetOrgByID(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	id := c.Params("id")

	filter := bson.D{{"id", id}}

	data := new(Model.ORG)
	err := col.FindOne(c.Context(), filter).Decode(&data)
	if err != nil {
		log.Println("GetOrgByParentID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetNameByID(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	id := c.Params("id")

	filter := bson.D{{"id", id}}

	data := new(Model.ORG)
	err := col.FindOne(c.Context(), filter).Decode(&data)
	if err != nil {
		log.Println("GetOrgByParentID Find ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data.Name)
}

func EditOrgName(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	var data struct {
		ID      string `json:"id"`
		NewName string `json:"new_name"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		log.Println("EditOrgName parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filter := bson.D{{"id", data.ID}}
	update := bson.D{{"$set", bson.D{{"name", data.NewName}}}}

	res, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println("EditOrgName UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if res.MatchedCount == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	result := new(Model.ORG)
	err = col.FindOne(c.Context(), bson.D{{"id", data.ID}}).Decode(&result)
	if err != nil {
		log.Println("EditOrgName FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteOrgById(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol
	data := new(Model.ORG)
	orgID := c.Params("id")

	filter := bson.D{{"id", orgID}}

	err := col.FindOne(c.Context(), filter).Decode(&data)
	if err != nil {
		log.Println("DeleteOrgById FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if data.Type == "team" {
		res, err := col.DeleteOne(c.Context(), filter)
		if res.DeletedCount > 1 {
			log.Println("Failed to delete")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err != nil {
			log.Println("DeleteOrgById DeleteOne ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

	} else if data.Type == "division" {
		if len(data.ChildrenID) > 0 {
			err := deleteChildArray(c, col, data.ChildrenID)
			if err != nil {
				log.Println("DeleteOrgById deleteChildArray ", err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		res, err := col.DeleteOne(c.Context(), filter)
		if res.DeletedCount > 1 {
			log.Println("Failed to delete")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err != nil {
			log.Println("DeleteOrgById DeleteOne ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if data.ParentID != "" {
		nFilter := bson.D{{"id", data.ParentID}}
		nUpdate := bson.D{{"$pull", bson.D{{"children_id", data.ID}}}}
		_, err := col.UpdateOne(c.Context(), nFilter, nUpdate)
		if err != nil {
			log.Println("DeleteOrgById delete parent ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.SendStatus(fiber.StatusOK)
}

func DeleteChildren(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	parentID := c.Params("id")

	_, err := col.UpdateOne(c.Context(), bson.D{{"id", parentID}}, bson.D{{"$set", bson.D{{"children_id", make([]string, 0)}}}})
	if err != nil {
		log.Println("DeleteChildren UpdateOne    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	org := new(Model.ORG)

	err = col.FindOne(c.Context(), bson.D{{"id", parentID}}).Decode(&org)
	if err != nil {
		log.Println("DeleteChildren FindOne    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

func deleteChildArray(c *fiber.Ctx, col *mongo.Collection, children []string) error {
	for k, id := range children {
		log.Println(k)
		filter := bson.D{{"id", id}}
		data := new(Model.ORG)
		err := col.FindOne(c.Context(), filter).Decode(&data)
		if err != nil {
			return err
		}
		if len(data.ChildrenID) > 0 {
			err := deleteChildArray(c, col, data.ChildrenID)
			if err != nil {
				return err
			}
		}
		_, err = col.DeleteOne(c.Context(), filter)
		if err != nil {
			return err
		}

		if data.ParentID != "" {
			nFilter := bson.D{{"id", data.ParentID}}
			nUpdate := bson.D{
				{"$pull", bson.D{{"children_id", data.ID}}},
			}
			_, err := col.UpdateOne(c.Context(), nFilter, nUpdate)
			if err != nil {
				log.Println("DeleteOrgById delete parent ", err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}

	return nil
}

func GetOrgStructByID(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	id := c.Params("id")

	var orgs []Model.ORG = make([]Model.ORG, 0)

	result, err := ReturnWholeorgStructUpward(col, id, orgs)
	if err != nil {
		log.Println("GetOrgStructByID    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func ReturnWholeorgStructUpward(col *mongo.Collection, id string, orgs []Model.ORG) ([]Model.ORG, error) {
	org := new(Model.ORG)
	err := col.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&org)
	if err != nil {
		log.Println("ReturnWholeorgStruct ", "id     ", err)
		return nil, err
	}

	if org.ParentID != "" {
		orgs, err = ReturnWholeorgStructUpward(col, org.ParentID, orgs)
		if err != nil {
			return nil, err
		}
	}

	orgs = append(orgs, *org)
	return orgs, nil
}

var orgParentSlice []string = make([]string, 0)
var orgChildSlice []string = make([]string, 0)

func GetOrgStructDownward(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	id := c.Params("parentID")

	orgParentSlice = nil
	orgParentSlice = append(orgParentSlice, id)

	var orgs []Model.ORG = make([]Model.ORG, 0)

	orgs, err := ReturnWholeOrgStructDownward(col, orgs)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

func ReturnWholeOrgStructDownward(col *mongo.Collection, orgs []Model.ORG) ([]Model.ORG, error) {
	org := new(Model.ORG)

	for _, v := range orgParentSlice {
		err := col.FindOne(context.TODO(), bson.D{{"id", v}}).Decode(&org)
		if err != nil {
			log.Println("ReturnWholeOrgStructDownward FindOne    ", err)
			return nil, err
		}
		orgs = append(orgs, *org)

		if len(org.ChildrenID) != 0 {
			for _, v := range org.ChildrenID {
				orgChildSlice = append(orgChildSlice, v)
			}
		}
	}

	if len(orgChildSlice) != 0 {
		orgParentSlice = orgChildSlice
		orgChildSlice = nil
		orgs, _ = ReturnWholeOrgStructDownward(col, orgs)
	}

	return orgs, nil

}

func GetAllTeams(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	filter := bson.D{{"type", "team"}}

	var orgs []Model.ORG = make([]Model.ORG, 0)

	cursor, err := col.Find(c.Context(), filter)
	if err != nil {
		log.Println("GetAllTeams Find    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &orgs)
	if err != nil {
		log.Println("GetAllTeams All    ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

// func GetFamilyByID(c *fiber.Ctx) error {
// 	col := DB.MI.OrgDBCol

// 	id := c.Params("id")

// 	results, err := test(col, id)
// 	if err != nil {
// 		log.Println("GetFamilyByID Failed", err)
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(results)
// }

// type orgStruct struct {
// 	ID         string      `json:"id" bson:"id"`
// 	Type       string      `json:"type" bson:"type"`
// 	ChildrenID []orgStruct `json:"children_id" bson:"children_id"`
// 	ParentID   string      `json:"parent_id" bson:"parent_id"`
// 	Name       string      `json:"name" bson:"name"`
// }

// func test(col *mongo.Collection, id string) ([]orgStruct, error) {
// 	org := new(Model.ORG)

// 	err := col.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&org)
// 	if err != nil {
// 		return nil, err
// 	}

// 	data := new(orgStruct)
// 	res, _ := json.Marshal(org)
// 	json.Unmarshal(res, &data)
// 	results := make([]orgStruct, 0)

// 	if len(org.ChildrenID) != 0 {

// 		for _, v := range org.ChildrenID {
// 			result, err := test(col, v)
// 			if err != nil {
// 				return nil, err
// 			}
// 			data.ChildrenID = append(data.ChildrenID, result...)
// 			results = append(results, *data)
// 		}

// 	} else {
// 		results = append(results, *data)
// 	}

// 	return results, nil
// }
