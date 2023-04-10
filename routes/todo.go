package routes

import (
	"example.com/crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.AlltTodos)
}
