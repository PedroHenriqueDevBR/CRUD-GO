package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Lista de usuarios")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Usuario cadastrado com sucesso")
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Usuario " + c.Params("id") + " atualizado com sucesso")
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Usuario " + c.Params("id") + " removido com sucesso")
	})

	app.Listen(":3000")
}
