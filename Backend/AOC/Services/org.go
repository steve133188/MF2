package Services

import (
	"fmt"
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
	fmt.Println(data.Type)
	if data.Type != "team" || data.Type != "division" {
		c.SendStatus(fiber.StatusBadRequest)
	}

	if data.ChildrenID == nil {
		data.ChildrenID = make([]string, 0)
	}

	if data.ParentID != "" {
		nFilter := bson.D{{"id", data.ParentID}}
		nUpdate := bson.D{{"$push", bson.D{{"children_id", data.ID}}}}
		res, err := col.UpdateOne(c.Context(), nFilter, nUpdate)
		if err != nil {
			log.Println("CreateDivision UpdateOne ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		fmt.Println(res.ModifiedCount)
		fmt.Println(res.MatchedCount)

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
