package user

import "github.com/gofiber/fiber/v2"

// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Data body model.PostUpdateRequest true "User data"
// @Success 200 {array} model.UserUpdateResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /user/update [put]
func Update(c *fiber.Ctx) error {
	return c.SendString("Update")
}
