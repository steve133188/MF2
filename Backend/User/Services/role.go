package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllRoles(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol

	cursor, err := col.Find(c.Context(), bson.D{{}})
	if err != nil {
		log.Println("GetAllRoles find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var result []Model.Roles = make([]Model.Roles, 0)
	err = cursor.All(c.Context(), &result)
	if err != nil {
		log.Println("GetAllRoles all: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetRoleByName(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol
	data := new(Model.Roles)
	name := c.Params("name")

	err := col.FindOne(c.Context(), bson.D{{"name", name}}).Decode(&data)
	if err != nil {
		log.Println("GetRoleByName FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func GetRolesName(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol

	var data []struct {
		Name string `json:"name"`
	}

	cursor, err := col.Find(c.Context(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("GetRolesName Find: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = cursor.All(c.Context(), &data)
	if err != nil {
		log.Println("GetRolesName All: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer cursor.Close(c.Context())

	var name []string
	for _, v := range data {
		if v.Name != "" {
			name = append(name, v.Name)
		}
	}

	return c.Status(fiber.StatusOK).JSON(name)
}

func AddRole(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol
	data := new(Model.Roles)
	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddRole FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = col.InsertOne(c.Context(), data)
	if err != nil {
		log.Println("AddRole InsertOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(data)

}

func UpdateRoleByName(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol
	data := new(Model.Roles)
	name := c.Params("name")
	err := c.BodyParser(&data)
	if err != nil {
		log.Println("UpdateRoleByName parse: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	col.FindOneAndUpdate(c.Context(), bson.D{{"name", name}}, bson.D{{"$set", data}})

	err = col.FindOne(c.Context(), bson.D{{"name", data.Name}}).Decode(&data)
	if err != nil {
		log.Println("UpdateRoleByName FindOne: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteRoleByName(c *fiber.Ctx) error {
	col := DB.MI.RoleDBCol

	name := c.Params("name")

	col.FindOneAndDelete(c.Context(), bson.D{{"name", name}})
	return c.SendStatus(fiber.StatusOK)

}
