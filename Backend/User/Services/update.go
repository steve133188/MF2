package Services

import (
	"fmt"
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserByID(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// user.Date = time.Now()
	user.ID = c.Params("id")
	update := bson.D{{Key: "$set", Value: user}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "id", Value: c.Params("id")}}, update)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func UpdateUserByName(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// user.Date = time.Now()
	user.ID = c.Params("name")
	update := bson.D{{Key: "$set", Value: user}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: c.Params("name")}}, update)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
func ChangeUserStatus(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// user.Date = time.Now()
	// user.Status = c.Params("name")
	update := bson.M{"$set": bson.M{
		"status": user.Status,
	}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func UpdateUserRole(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.M{"$set": bson.M{
		"role":      user.Role,
		"authority": user.Authority,
	}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func ChangeUserPassword(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	password, err := Util.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert agent",
			"error":   err,
		})
	}
	update := bson.M{"$set": bson.M{"password": password}}

	err = usersCollection.FindOneAndUpdate(c.Context(), bson.D{{Key: "email", Value: user.Email}}, update).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

func UpdateDivisionAndTeam(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.M{"$set": bson.M{"division_name": user.DivisionName, "team": user.Team}}

	err := usersCollection.FindOneAndUpdate(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}
