package hanlder

import "github.com/gofiber/fiber/v2"

func UpdateFlow(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}

func UpdateAction(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}

func UpdateOption(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}
