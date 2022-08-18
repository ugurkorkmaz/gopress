package auth

import "github.com/gofiber/fiber/v2"

// @Summary Login
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce application/json
// @Param Data body model.LoginRequest true "Login Request"
// @Success 200 {array} model.LoginResponse
// @Failure 400 {array} model.Exception
// @Router /auth/login [get] [post]
func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}
