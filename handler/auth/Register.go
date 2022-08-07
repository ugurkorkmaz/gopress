package auth

import "github.com/gofiber/fiber/v2"

// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce application/json
// @Success 200 {array} ent.User
// @Failure 400 {json} string
// @Failure 500 {json} string
// @Router /api/auth/register [post]
func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}
