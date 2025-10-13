package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID 			int `json:"id"`
	Completed 	bool `json:"completed"`
	Body 		string `json:"body"`
}

func main() {
	fmt.Println("Hello World this is vanya")
	app := fiber.New()

	// todos array
	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"hello world"})
	})

	// CREATE TODO
	// adding new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// {id:0, completed:false, body:""}
		// memory address of todo
		todo := &Todo{}

		// body parser untuk parsing json ke struct
		// body parser binds the request body to the struct
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg":"body is required"})
		}

		todo.ID = len(todos) + 1
		// getting the value *todo
		todos = append(todos, *todo)


		return c.Status(201).JSON(todo)
	})

	// UPDATE TODO BY ID
	// For checking a todo as completed
	app.Patch("/api/todos/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})

	// DELETE TODO BY ID
	// deleting a todo
	app.Delete("/api/todos/:id")

	log.Fatal(app.Listen(":4000"))
}