package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Return all users
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Lista de usuarios")
	})

	// Validate and save user
	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Usuario cadastrado com sucesso")
	})

	// search and update user
	app.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Usuario " + c.Params("id") + " atualizado com sucesso")
	})

	// search and delete user
	app.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Usuario " + c.Params("id") + " removido com sucesso")
	})

	app.Listen(":3000")
}
