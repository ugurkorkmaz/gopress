package user

import "github.com/gofiber/fiber/v2"

// @Summary Show user
// @Description Show user
// @Tags User
// @Accept json
// @Produce application/json
// @Param Id path string true "User Id"
// @Success 200 {array} model.UserShowResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /user/show/{id} [get]
func Show(c *fiber.Ctx) error {
	return c.SendString("Show")
}
