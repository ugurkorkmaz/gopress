package post

import "github.com/gofiber/fiber/v2"

// @Summary Create a post
// @Description Create a post
// @Tags Post
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Data body model.PostCreateRequest true "Post data"
// @Success 200 {array} model.PostCreateResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /post/create [post]
func Create(c *fiber.Ctx) error {
	return c.SendString("Create")
}
