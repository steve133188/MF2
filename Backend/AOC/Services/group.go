package Services

import (
	"mf-aoc-service/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddGroup(c *fiber.Ctx) error {
	customersCollection := DB.MI.GrpDBCol

	var data struct {
		Name []string `json:"name" bson:"name"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	update := bson.D{{"$addToSet", bson.D{{"group", bson.D{{"$each", data.Name}}}}}}
	result, err := customersCollection.UpdateOne(c.Context(), bson.D{{}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})

}

// 	data.ID = xid.New().String()

// 	result, err := customersCollection.InsertOne(c.Context(), data)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Failed to insert",
// 			"error":   err,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    result,
// 	})
// }

// func EditGroup(c *fiber.Ctx) error {
// 	col := DB.MI.GrpDBCol

// 	data := new(Model.EditGroup)
// 	err := c.BodyParser(&data)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 			"error":   err,
// 		})
// 	}

// 	filter := bson.D{{"name", data.Old}}
// 	update := bson.D{{"name", data.New}}

// 	result, err := col.UpdateOne(c.Context(), filter, update)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Failed to update",
// 			"error":   err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 	})
// }
