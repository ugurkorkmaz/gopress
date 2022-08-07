package auth

import "github.com/gofiber/fiber/v2"

func Guard(c *fiber.Ctx) error {
	return c.Next()
}
