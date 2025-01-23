package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "Hello, World!",
			"timestamp": time.Now().Unix(),
		})
	})

	app.Listen(":8080")
}
