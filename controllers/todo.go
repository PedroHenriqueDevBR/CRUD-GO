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
			"message": "Error: Verify JS argumentON",
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
			"message": "Error: Verify Id argument",
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

func UpdateTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error: Verify id argument",
		})
	}

	type Request struct {
		Title *string `json:"title"`
		Done  *bool   `json:"done"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error: Verify JS argumentON",
		})
	}

	var todo *Todo

	for _, t := range savedTodos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Done != nil {
		todo.Done = *body.Done
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error: Verify id argument",
		})
	}

	for i, todo := range savedTodos {
		if todo.Id == id {
			savedTodos = append(savedTodos[:i], savedTodos[i+1:]...)

			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"success": true,
				"message": "Deleted Succesfully",
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
