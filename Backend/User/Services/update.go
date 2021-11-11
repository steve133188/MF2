package Services

import (
	"log"
	"mf-user-servies/DB"
	"mf-user-servies/Model"
	"mf-user-servies/Util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserByName(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)

	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}
	password, err := Util.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to hass password",
			"error":   err.Error(),
		})
	}

	user.Password = password

	update := bson.D{{Key: "$set", Value: user}}

	_, err = usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func ChangeUserStatus(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)

	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	update := bson.M{"$set": bson.M{
		"status": user.Status,
	}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

func UpdateUserRole(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
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

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

func UpdateDivisionAndTeam(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	update := bson.M{"$set": bson.M{"division_name": user.DivisionName, "team": user.Team}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteUserTeam(c *fiber.Ctx) error {
	usersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(Model.User)
	data := new(Model.User)
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	update := bson.M{"$set": bson.M{"division_name": "", "team": ""}}

	_, err := usersCollection.UpdateOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}

	err = usersCollection.FindOne(c.Context(), bson.D{{Key: "username", Value: user.UserName}}).Decode(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func UpdateChannelInfoByID(c *fiber.Ctx) error {
	customersCollection := DB.MI.DBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// customer := new(Model.Customer)
	chanInfo := new(Model.Info)

	if err := c.BodyParser(&chanInfo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.D{{"$set", bson.D{{"channel_info", chanInfo}}}}

	_, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "phone", Value: chanInfo.Phone}}, update)
	// fmt.Println(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Customer failed to update",
			"error":   err.Error(),
		})
	}
	customer := new(Model.User)
	customersCollection.FindOne(c.Context(), bson.D{{"phone", chanInfo.Phone}}).Decode(&customer)
	return c.Status(fiber.StatusCreated).JSON(chanInfo)
}
