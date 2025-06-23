package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Get("/:name?", func(a *fiber.Ctx) error {
		if a.Params("name") != "" {
			return a.SendString("Hello" + a.Params("name"))
		}
		return a.SendString("Where is John?")
	})

	log.Fatal(app.Listen(":8080"))
}
