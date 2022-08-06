package main

import (
	"gopress/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.Index)

	app.Listen(":3000")
}
