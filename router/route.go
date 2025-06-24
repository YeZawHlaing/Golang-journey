package main

import "github.com/gofiber/fiber/v2"

func main() {

	router()

}

func router() {
	app := fiber.New()

	app.Route("/product", func(endpoint fiber.Router) {

		endpoint.Get("/allproduct", func(c *fiber.Ctx) error {
			return c.SendString("get product")
		})
		endpoint.Post("/upload", func(c *fiber.Ctx) error {
			return c.SendString("upload product")
		})
		endpoint.Put("/:id", func(c *fiber.Ctx) error {
			return c.SendString("update product")
		})
		endpoint.Delete("/:id", func(c *fiber.Ctx) error {
			return c.SendString("delete product")
		})

	})
	app.Listen(":3000")
}
