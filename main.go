package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	return app
}

func main() {
	fmt.Println("Hello, world!")
	app := setupApp()
	app.Listen(":80")
}
