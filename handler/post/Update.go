package post

import "github.com/gofiber/fiber/v2"

// @Summary Update a post
// @Description Update a post
// @Tags Post
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Data body model.PostUpdateRequest true "Post data"
// @Success 200 {array} model.PostUpdateResponse "ok"
// @Failure 400 {array} model.Exception
// @Router /post/update [put]
func Update(c *fiber.Ctx) error {
	return c.SendString("Update")
}
