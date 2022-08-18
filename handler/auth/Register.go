package auth

import "github.com/gofiber/fiber/v2"

// @Summary Register
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce application/json
// @Param Data body model.RegisterRequest true "Register Request"
// @Success 200 {array} model.RegisterResponse
// @Failure 400 {array} model.Exception
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}
