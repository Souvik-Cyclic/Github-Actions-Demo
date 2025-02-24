package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":  "Hello, World!",
			"message1": "You are welcome to the world of Go, Have a nice day!",
			"message2": "This is a simple Go application running in a Docker container",
			"timestamp": time.Now().Unix(),
		})
	})

	app.Listen(":8080")
}
