package Services

import (
	"mf-aoc-service/DB"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteChannelById(c *fiber.Ctx) error {
	userCollection := DB.MI.ChanDBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.M{"id": paramID}

	err := userCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Delete (ID = " + paramID + ") success",
	})
}

func DeleteRoleById(c *fiber.Ctx) error {
	collection := DB.MI.RoleDBCol

	// get param
	paramID := c.Params("id")

	// find and delete todo
	query := bson.D{{Key: "id", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func DeleteRoleByName(c *fiber.Ctx) error {
	collection := DB.MI.RoleDBCol

	// get param
	paramID := c.Params("name")

	// find and delete todo
	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func DeleteTagsByName(c *fiber.Ctx) error {
	collection := DB.MI.TagsDBCol

	// get param
	paramID := c.Params("name")

	// find and delete todo
	query := bson.D{{Key: "name", Value: paramID}}

	err := collection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func DeleteOrganizationByPhone(c *fiber.Ctx) error {
	userCollection := DB.MI.OrgDBCol

	// get param
	paramID := c.Params("phone")

	// find and delete todo
	query := bson.M{"phone": paramID}

	err := userCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Delete (ID = " + paramID + ") success",
	})
}
