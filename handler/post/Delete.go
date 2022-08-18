package post

import "github.com/gofiber/fiber/v2"

// @Summary Delete a post
// @Description Delete a post
// @Tags Post
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Post Id path string true "Post Id"
// @Success 200 {array} model.PostDeleteResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /post/delete/{id} [delete]
func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
