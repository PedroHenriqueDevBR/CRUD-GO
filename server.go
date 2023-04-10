package main

import (
	"example.com/crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Working",
		})
	})

	api := app.Group("/api")
	routes.TodoRoute(api.Group("/todo"))

}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(":3000")
}
