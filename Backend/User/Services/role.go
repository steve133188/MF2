package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddRoleToUser(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	data := new(Model.Roles)
	err := c.BodyParser(&data)
	if err != nil {
		log.Println("AddRoleToUser parse ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	phone := data.Phone

	filter := bson.D{{"phone", phone}}
	user := new(Model.User)
	err = col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("AddRoleToUser FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if user.Role != "" {
		log.Println("AddRoleToUser User existed role", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user.Role = data.Role
	user.Authority = data.Auth

	res, err := col.UpdateOne(c.Context(), filter, &user)
	if err != nil {
		log.Println("AddRoleToUser UpdateOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.ModifiedCount == 0 {
		log.Println("AddRoleToUser Update failed ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUserRole(c *fiber.Ctx) error {
	col := DB.MI.UserDBCol

	user := new(Model.User)
	phone := c.Params("phone")

	auth := Model.Auth{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	filter := bson.D{{"phone", phone}}

	err := col.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		log.Println("DeleteUserRole FindOne ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.Role = ""
	user.Authority = auth
	// user.Authority.Admin = false
	// user.Authority.Boardcast = false
	// user.Authority.Contact = false
	// user.Authority.Dashboard = false
	// user.Authority.Flowbuilder = false
	// user.Authority.Integrations = false
	// user.Authority.Livechat = false
	// user.Authority.Organization = false
	// user.Authority.ProductCatalogue = false

}
