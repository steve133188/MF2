package hanlder

import "github.com/gofiber/fiber/v2"

func GetCFlowByFlowID(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}

func GetCFlowByCompany(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}

func GetCActionByActionID(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("PONG")
}
