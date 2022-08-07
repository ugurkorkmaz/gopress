package main

import (
	"gopress/app/database"
	"gopress/app/environment"
	"gopress/handler/auth"
	"gopress/handler/post"
	"gopress/handler/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "gopress/docs"
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

	server.Use(logger.New(), cors.New())
	server.Get("/swagger/*", swagger.HandlerDefault)
}

// @title GoPress API Documentation
// @version 0.0.1
// @description GoPress API Documentation
// @termsOfService http://github.com/ugurkorkmaz/gopress/
// @contact.name API Support
// @contact.email ugur@extends.work
// @license.name MIT
// @license.url http://github.com/ugurkorkmaz/gopress
// @host localhost:3000
// @BasePath /
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
