package Services

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"mf-auth-servies/DB"
	"mf-auth-servies/Model"
)

func Login (c * fiber.Ctx) error{

	//e := c.Params("email")
	//pwd :=c.Params("password")

	//if (e != "stevechakcy@gmail.com" ) || (pwd != "1234"){
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"success": true,
	//		"message":"email or password wrong",
	//	})
	//}

	//postBody, _ := json.Marshal(map[string]string{
	//	"password":  "1234",
	//	"email": "steve@example.com",
	//})
	//responseBody := bytes.NewBuffer(postBody)
	//resp, err :=http.Post("http//:localhost:3000/login" ,"application/json" , responseBody)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	users_collection := DB.MI.DBCol
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user Model.User

	q:= &Model.UserCredential{
		Email: c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	//if err := c.BodyParser(&q); err != nil {
	//	return err
	//}

	fmt.Println(q)

	filter := bson.M{"email":q.Email,"password":q.Password}

	findResult :=users_collection.FindOne(ctx , filter)
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	err := findResult.Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//sb := string(body)
	//log.Printf(sb)
	validToken, err := GenJWT()
	fmt.Println(validToken)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Please Login", // invalid token
			"error":   err,
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message":"authenticated",
		"data": fiber.Map{
			"token":string(validToken),
		},
	})
}