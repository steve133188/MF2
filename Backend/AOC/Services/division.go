package Services

import (
	"fmt"
	"mf-aoc-service/DB"
	"mf-aoc-service/Model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateDivision(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Division)
	exist := new(Model.Division)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	isExisted := customersCollection.FindOne(c.Context(), bson.D{{Key: "name", Value: data.Name}}).Decode(exist)
	fmt.Println(isExisted)
	if (isExisted) == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Is exist",
		})
	}

	data.CreatedAt = time.Now().String()
	result, err := customersCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	// get the inserted data
	customer := &Model.Division{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	customersCollection.FindOne(c.Context(), query).Decode(customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

func GetDivisionByName(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	data := new(Model.Division)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	filter := bson.D{{"name", data.Name}}
	err = col.FindOne(c.Context(), filter).Decode(&data)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})

}

func GetAllDivision(c *fiber.Ctx) error {
	collection := DB.MI.OrgDBCol
	query := bson.D{{}}

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find context",
			"error":   err.Error(),
		})
	}

	var todos []Model.Division = make([]Model.Division, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error to interate cursor into result",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(c.Context())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    todos,
	})
}

func UpdateDivisionByName(c *fiber.Ctx) error {
	collection := DB.MI.OrgDBCol
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(Model.Division)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	update := bson.D{{Key: "$set", Value: todo}}

	_, err := collection.UpdateOne(c.Context(), bson.D{{Key: "name", Value: c.Params("name")}}, update)
	fmt.Println(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func CreateTeam(c *fiber.Ctx) error {
	customersCollection := DB.MI.OrgDBCol

	data := new(Model.Division)
	exist := new(Model.Division)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	isExisted := customersCollection.FindOne(c.Context(), bson.D{{Key: "name", Value: data.Name}}).Decode(exist)
	// var val bson.A
	if (isExisted) == nil { //existing division

		update := bson.D{{"$addToSet", bson.D{{"team", bson.D{{"$each", data.Team}}}}}}
		result, err := customersCollection.UpdateOne(c.Context(), bson.D{{Key: "name", Value: data.Name}}, update)
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

	data.ID = xid.New().String()

	result, err := customersCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func UpdateTeam(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	data := new(Model.EditTeam)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse",
			"error":   err,
		})
	}

	filter := bson.D{{"name", data.DivName}, {"team", data.Old}}
	update := bson.D{{"$set", bson.D{{"team.$", data.New}}}}

	result, err := col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "Team not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Update Success",
	})

}

func DelTeam(c *fiber.Ctx) error {
	col := DB.MI.OrgDBCol

	data := new(Model.EditTeam)

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse",
			"error":   err,
		})
	}

	filter := bson.D{{"name", data.DivName}}
	update := bson.D{{"$pull", bson.D{{"team", data.Old}}}}
	_, err = col.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Update Success",
	})
}
