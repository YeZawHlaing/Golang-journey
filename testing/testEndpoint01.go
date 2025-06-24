package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	testing()
}

func testing() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Static("/", "./application")

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber API!")
	})

	log.Fatal(app.Listen(":8082"))
}
