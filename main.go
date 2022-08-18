package main

import (
	"gopress/app/database"
	"gopress/app/environment"
	"gopress/handler/auth"
	"gopress/handler/exception"
	"gopress/handler/post"
	"gopress/handler/user"
	"gopress/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
	server = fiber.New(fiber.Config{
		AppName:      "GoPress",
		ServerHeader: "GoPress",
		ErrorHandler: exception.Throw,
	})

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
// @BasePath /api
func main() {
	api := server.Group("/api")

	auths := api.Group("/auth").Name("auth.")
	auths.Post("/register", auth.Register).Name("register")
	auths.Post("/login", auth.Login).Name("login")

	posts := api.Group("/post").Name("post.")
	posts.Get("/show/:id", post.Show).Name("show")
	posts.Get("/list", post.List).Name("list")
	posts.Use(auth.Guard).Post("/create", post.Create).Name("create")
	posts.Use(auth.Guard).Put("/update", post.Update).Name("update")
	posts.Use(auth.Guard).Delete("/delete/:id", post.Delete).Name("delete")

	users := api.Group("/user").Name("user.")
	users.Get("/:id", user.Show).Name("show")
	users.Use(auth.Guard).Put("/:id", user.Update).Name("update")
	users.Use(auth.Guard).Delete("/:id", user.Delete).Name("delete")
	users.Use(auth.Guard).Get("/", user.List).Name("list")

	server.Use("/", filesystem.New(filesystem.Config{
		Root:         template.Dist(),
		NotFoundFile: "index.html",
	}))
	server.Listen(":3000")
}
