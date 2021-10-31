package Services

// func GetAllOrganization(c *fiber.Ctx) error {
// 	fmt.Println("getall")
// 	// token := c.Request().Header.Peek("Authorization")
// 	// _, err := Util.ParseToken(string(token))
// 	// if err != nil {
// 	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 	// 		"error": "Unauthorized",
// 	// 	})
// 	// }

// 	usersCollection := DB.MI.DBCol
// 	// Query to filter
// 	query := bson.D{{}}

// 	cursor, err := usersCollection.Find(c.Context(), query)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Failed to find context",
// 			"error":   err.Error(),
// 		})
// 	}

// 	var users []Model.Organization = make([]Model.Organization, 0)

// 	// iterate the cursor and decode each item into a Todo
// 	err = cursor.All(c.Context(), &users)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Error to interate cursor into result",
// 			"error":   err.Error(),
// 		})
// 	}

// 	// for i := range users {
// 	// 	users[i].Date = users[i].Date.Add(time.Hour * 8)
// 	// }

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    users,
// 	})
// }
