package exception

import "github.com/gofiber/fiber/v2"

func Throw(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: application/json; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Return statuscode with error message
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
		"code":    code,
	})
}
