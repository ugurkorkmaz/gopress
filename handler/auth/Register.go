package auth

import "github.com/gofiber/fiber/v2"

// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce application/json
// @Success 200 {array} entity.User
// @Failure 400 {array} entity.Failure
// @Router /api/auth/register [post]
func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}
