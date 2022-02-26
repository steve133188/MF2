package hanlder

import "github.com/gofiber/fiber/v2"

func CreateFlow(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}

func CreateOption(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}

func CreateAction(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}
