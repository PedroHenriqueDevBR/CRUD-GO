package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var savedTodos = []*Todo{
	{
		Id:    1,
		Title: "Read the book",
		Done:  false,
	},
	{
		Id:    2,
		Title: "Run",
		Done:  false,
	},
}

func AlltTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": savedTodos,
		},
	})
}
