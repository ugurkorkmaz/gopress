package main

import (
	"gopress/app/database"
	"gopress/app/environment"
	"gopress/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	if err := environment.Load(); err != nil {
		panic(err)
	}
	if err := database.Connect(); err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/", handler.Index)

	app.Listen(":3000")
}
