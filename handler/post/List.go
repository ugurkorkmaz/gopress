package post

import "github.com/gofiber/fiber/v2"

// @Summary list posts
// @Description list posts
// @Tags Post
// @Accept json
// @Produce application/json
// @Param page query int false "Page" default(1) minimum(1) maximum(10)
// @Param limit query int false "Limit" default(10) minimum(1) maximum(10)
// @Success 200 {array} model.PostListResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /post/list [get]
func List(c *fiber.Ctx) error {
	return c.SendString("List")
}
