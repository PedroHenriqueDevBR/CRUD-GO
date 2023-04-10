package routes

import (
	"example.com/crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.AlltTodos)
	route.Post("", controllers.CreateTodo)
	route.Get("/:id", controllers.GetTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
}
