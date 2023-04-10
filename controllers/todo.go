package controllers

import (
	"fmt"
	"strconv"

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

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	todo := &Todo{
		Id:    len(savedTodos) + 1,
		Title: body.Title,
		Done:  false,
	}

	savedTodos = append(savedTodos, todo)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}

	for _, todo := range savedTodos {
		if todo.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Not found",
	})
}
