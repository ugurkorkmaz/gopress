package user

import "github.com/gofiber/fiber/v2"

func Update(c *fiber.Ctx) error {
	return c.SendString("Update")
}
