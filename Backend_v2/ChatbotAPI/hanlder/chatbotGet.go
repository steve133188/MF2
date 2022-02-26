package hanlder

import "github.com/gofiber/fiber/v2"

func GetCFlowByFlowID(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}

func GetCFlowByCompany(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}

func GetCActionByActionID(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON()
}
