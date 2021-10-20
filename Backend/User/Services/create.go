// package Services

// import (
// 	"context"
// 	"github.com/gofiber/fiber/v2"
// 	"mf-user-servies/DB"
// 	"mf-user-servies/Model"
// 	"time"
// )

//type CreateUser interface {
//	CreateOneUser(obj interface{}) (status string,err error)
//	CreateManyUser(obj interface{}) (status string,err error)
//}
//
//type Create struct {}
//
//func NewCreate() *Create{
//	return &Create{}
//}
//
//func CreateOneUser(obj interface{}) (status string,err error){
//	return
//}
//
//func CreateManyUser(obj interface{}) (status string,err error){
//	return
//}

// func CreateOneUserTest(c * fiber.Ctx) error{
// 	col := DB.MI.DBCol
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	//user := Model.User{
// 	//	Username: "steve",
// 	//	Password: "123456",
// 	//}
// 	user :=new(Model.User)

// 	//if err := c.BodyParser(user); err != nil {
// 	//	log.Println(err)
// 	//	return c.Status(400).JSON(fiber.Map{
// 	//		"success": false,
// 	//		"message": "Failed to parse body",
// 	//		"error":   err,
// 	//	})
// 	//}  No need to fix this

// 	result, err := col.InsertOne(ctx, user)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Catchphrase failed to insert",
// 			"error":   err,
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"data":    result,
// 		"success": true,
// 		"message": "Catchphrase inserted successfully",
// 	})

// }

package Services

import (
	"fmt"
	"mf-user-servies/DB"
	"mf-user-servies/Model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol

	data := new(Model.User)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	result, err := usersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// get the inserted data
	user := &Model.User{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	fmt.Println(result.InsertedID)
	usersCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}
