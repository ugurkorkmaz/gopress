package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Guard(c *fiber.Ctx) error {
	auth := true

	if auth {
		return c.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized, "not authorized")
}
