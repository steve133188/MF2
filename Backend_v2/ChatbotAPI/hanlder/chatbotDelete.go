package hanlder

import "github.com/gofiber/fiber/v2"

func DeleteFlow(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusOK)
}

func DeleteOption(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusOK)
}

func DeleteAction(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusOK)
}
