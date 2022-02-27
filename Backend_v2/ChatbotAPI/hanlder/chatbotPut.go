package hanlder

import "github.com/gofiber/fiber/v2"

func UpdateFlow(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}

func UpdateAction(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}

func UpdateOption(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}
