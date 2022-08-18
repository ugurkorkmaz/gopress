package post

import "github.com/gofiber/fiber/v2"

// @Summary Post show
// @Description Post show
// @Tags Post
// @Accept json
// @Produce application/json
// @Param Id path string true "Post Id"
// @Success 200 {array} model.PostDeleteResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /post/show/{id} [get]
func Show(c *fiber.Ctx) error {
	return c.SendString("Show")
}
