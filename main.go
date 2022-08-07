package main

import (
	"gopress/app/database"
	"gopress/app/environment"
	"gopress/handler/auth"
	"gopress/handler/post"
	"gopress/handler/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var server *fiber.App

func init() {
	if err := environment.Load(); err != nil {
		panic(err)
	}
	if err := database.Connect(); err != nil {
		panic(err)
	}
	server = fiber.New()

	server.Use(logger.New())
}

func main() {
	api := server.Group("/api")

	auths := api.Group("/auth")
	auths.Post("/register", auth.Register)
	auths.Post("/login", auth.Login)

	posts := api.Group("/post", auth.Guard)
	posts.Post("/", post.Create)
	posts.Get("/:id", post.Show)
	posts.Put("/:id", post.Update)
	posts.Delete("/:id", post.Delete)
	posts.Get("/", post.List)

	users := api.Group("/user", auth.Guard)
	users.Get("/:id", user.Show)
	users.Put("/:id", user.Update)
	users.Delete("/:id", user.Delete)
	users.Get("/", user.List)

	server.Listen(":3000")
}
